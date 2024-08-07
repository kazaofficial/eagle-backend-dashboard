package service

import (
	"context"
	"eagle-backend-dashboard/dto"
	"eagle-backend-dashboard/entity"
	"eagle-backend-dashboard/repository"
	"log"
	"os"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
)

type OlahDataServiceImpl struct {
	daftarProsesPenarikanDataRepository repository.DaftarProsesPenarikanDataRepository
}

func NewOlahDataService(daftarProsesPenarikanDataRepository repository.DaftarProsesPenarikanDataRepository) OlahDataService {
	return &OlahDataServiceImpl{
		daftarProsesPenarikanDataRepository: daftarProsesPenarikanDataRepository,
	}
}

func SSHToServer() (*ssh.Client, error) {
	host := os.Getenv("OLAHDATA_HOST")
	username := os.Getenv("OLAHDATA_USERNAME")
	privateFilePath := os.Getenv("OLAHDATA_PRIVATE_FILE_PATH")

	key, err := os.ReadFile(privateFilePath)
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
		return nil, err
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
		return nil, err
	}

	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", host, config)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
		return nil, err
	}

	return client, nil
}

func GeneratePenarikanData(request *dto.DaftarProsesPenarikanDataRequest) error {
	client, err := SSHToServer()
	if err != nil {
		return err
	}

	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("failed to create session: %v", err)
		return err
	}
	defer session.Close()
	log.Println("Connected to server")

	var b []byte
	output := "python3 generate_ingestion_raw_dag_" + request.ConnType +
		".py " + request.SourceConnectionID +
		" " + request.TargetConnectionID +
		" " + request.SchemaTable +
		" '@" + request.Schedule +
		"'"
	log.Println(output)
	if b, err = session.Output(output); err != nil {
		log.Fatalf("failed to run: %v", err)
		return err
	}

	stringResult := string(b)
	log.Println(stringResult)

	if !strings.Contains(stringResult, "DAG file created") {
		return errors.New(stringResult)
	}

	return nil
}

func (service *OlahDataServiceImpl) TestSSHToServer() error {
	client, err := SSHToServer()
	if err != nil {
		return err
	}

	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("failed to create session: %v", err)
		return err
	}
	defer session.Close()
	log.Println("Connected to server")

	var b []byte
	if b, err = session.Output("echo \"Hello World\""); err != nil {
		log.Fatalf("failed to run: %v", err)
		return err
	}

	log.Println(string(b))

	return nil
}

func (service *OlahDataServiceImpl) GetDaftarProsesPenarikanData(ctx context.Context, request *dto.DaftarProsesPenarikanDataListRequest) ([]dto.DaftarProsesPenarikanDataResponse, *dto.Pagination, error) {
	offset := 0
	page := 1
	limit := 10

	if request.Page != nil {
		page = *request.Page
		offset = (page - 1) * limit
	}

	if request.Limit != nil {
		limit = *request.Limit
	}

	sort := "id desc"
	if request.Sort != "" {
		sort = request.Sort
		sort = strings.ReplaceAll(sort, ".", " ")
	}

	daftarProsesPenarikanData, err := service.daftarProsesPenarikanDataRepository.GetDaftarProsesPenarikanData(ctx, &limit, &offset, &sort, request.Search)
	if err != nil {
		return nil, nil, err
	}

	countDaftarProsesPenarikanDatas, err := service.daftarProsesPenarikanDataRepository.CountDaftarProsesPenarikanData(ctx, request.Search)
	if err != nil {
		return nil, nil, err
	}

	daftarProsesPenarikanDataResponses := []dto.DaftarProsesPenarikanDataResponse{}
	for _, daftarProsesPenarikanData := range daftarProsesPenarikanData {
		daftarProsesPenarikanDataResponses = append(daftarProsesPenarikanDataResponses, ConvertDaftarProsesPenarikanDataEntityToDTO(daftarProsesPenarikanData))
	}

	pagination := dto.Pagination{
		Page:      page,
		Limit:     limit,
		Total:     len(daftarProsesPenarikanData),
		TotalData: countDaftarProsesPenarikanDatas,
		TotalPage: countDaftarProsesPenarikanDatas/limit + 1,
	}

	return daftarProsesPenarikanDataResponses, &pagination, nil
}

func (service *OlahDataServiceImpl) GetDaftarProsesPenarikanDataByID(ctx context.Context, id int) (*dto.DaftarProsesPenarikanDataResponse, error) {
	daftarProsesPenarikanData, err := service.daftarProsesPenarikanDataRepository.GetDaftarProsesPenarikanDataByID(ctx, id)
	if err != nil {
		return nil, err
	}

	daftarProsesPenarikanDataResponse := ConvertDaftarProsesPenarikanDataEntityToDTO(*daftarProsesPenarikanData)
	return &daftarProsesPenarikanDataResponse, nil
}

func (service *OlahDataServiceImpl) CreateDaftarProsesPenarikanData(ctx context.Context, request *dto.DaftarProsesPenarikanDataRequest) (*dto.DaftarProsesPenarikanDataResponse, error) {
	err := GeneratePenarikanData(request)
	if err != nil {
		return nil, err
	}
	daftarProsesPenarikanData := entity.DaftarProsesPenarikanData{
		ConnType:           request.ConnType,
		SourceConnectionID: request.SourceConnectionID,
		TargetConnectionID: request.TargetConnectionID,
		SchemaTable:        request.SchemaTable,
		Schedule:           request.Schedule,
	}

	err = service.daftarProsesPenarikanDataRepository.CreateDaftarProsesPenarikanData(ctx, &daftarProsesPenarikanData)
	if err != nil {
		return nil, err
	}

	daftarProsesPenarikanDataResponse := ConvertDaftarProsesPenarikanDataEntityToDTO(daftarProsesPenarikanData)
	return &daftarProsesPenarikanDataResponse, nil
}

func (service *OlahDataServiceImpl) UpdateDaftarProsesPenarikanData(ctx context.Context, id int, request *dto.DaftarProsesPenarikanDataRequest) (*dto.DaftarProsesPenarikanDataResponse, error) {
	err := GeneratePenarikanData(request)
	if err != nil {
		return nil, err
	}

	daftarProsesPenarikanData, err := service.daftarProsesPenarikanDataRepository.GetDaftarProsesPenarikanDataByID(ctx, id)
	if err != nil {
		return nil, err
	}

	daftarProsesPenarikanData.ConnType = request.ConnType
	daftarProsesPenarikanData.SourceConnectionID = request.SourceConnectionID
	daftarProsesPenarikanData.TargetConnectionID = request.TargetConnectionID
	daftarProsesPenarikanData.SchemaTable = request.SchemaTable
	daftarProsesPenarikanData.Schedule = request.Schedule

	err = service.daftarProsesPenarikanDataRepository.UpdateDaftarProsesPenarikanData(ctx, daftarProsesPenarikanData)
	if err != nil {
		return nil, err
	}

	daftarProsesPenarikanDataResponse := ConvertDaftarProsesPenarikanDataEntityToDTO(*daftarProsesPenarikanData)
	return &daftarProsesPenarikanDataResponse, nil
}

func (service *OlahDataServiceImpl) DeleteDaftarProsesPenarikanData(ctx context.Context, id int) (*dto.DaftarProsesPenarikanDataResponse, error) {
	daftarProsesPenarikanData, err := service.daftarProsesPenarikanDataRepository.GetDaftarProsesPenarikanDataByID(ctx, id)
	if err != nil {
		return nil, err
	}

	err = service.daftarProsesPenarikanDataRepository.DeleteDaftarProsesPenarikanData(ctx, id)
	if err != nil {
		return nil, err
	}

	daftarProsesPenarikanDataResponse := ConvertDaftarProsesPenarikanDataEntityToDTO(*daftarProsesPenarikanData)
	return &daftarProsesPenarikanDataResponse, nil
}

func ConvertDaftarProsesPenarikanDataEntityToDTO(daftarProsesPenarikanData entity.DaftarProsesPenarikanData) dto.DaftarProsesPenarikanDataResponse {
	daftarProsesPenarikanDataResponse := dto.DaftarProsesPenarikanDataResponse{
		ID:                 daftarProsesPenarikanData.ID,
		ConnType:           daftarProsesPenarikanData.ConnType,
		SourceConnectionID: daftarProsesPenarikanData.SourceConnectionID,
		TargetConnectionID: daftarProsesPenarikanData.TargetConnectionID,
		SchemaTable:        daftarProsesPenarikanData.SchemaTable,
		Schedule:           daftarProsesPenarikanData.Schedule,
		CreatedAt:          daftarProsesPenarikanData.CreatedAt,
		UpdatedAt:          daftarProsesPenarikanData.UpdatedAt,
		DeletedAt:          daftarProsesPenarikanData.DeletedAt,
	}

	return daftarProsesPenarikanDataResponse
}
