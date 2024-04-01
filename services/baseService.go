package services

import (
	"encoding/json"
	"log"
	"notepad/config"
	"notepad/models"
	"notepad/utils"
	"os"
	"path"

	"github.com/google/uuid"
)

type ServiceResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
}

func success[T any](data T) ServiceResponse {
	return ServiceResponse{
		Code:    200,
		Message: "",
		Data:    data,
		Success: true,
	}
}

func failure[T any](code int, message string, data T) ServiceResponse {
	log.Println(code, message)
	return ServiceResponse{
		Code:    200,
		Message: "",
		Data:    data,
		Success: true,
	}
}

type BaseService[T models.BaseModel] struct {
	name     string
	dataPath string
	baseList []T
}

type DataFile[T any] struct {
	List []T `json:"list"`
}

func NewBaseService[T models.BaseModel](name string) *BaseService[T] {
	s := &BaseService[T]{
		name: name,
	}
	s.dataPath = path.Join(config.Config.BaseDataPath, name)
	s.readFile()
	return s
}

func (s *BaseService[T]) readFile() {
	// todo 后续换成读取接口，现在先直接读本地文件

	_, err := os.Stat(s.dataPath)
	if err != nil {
		if os.IsExist((err)) {
			log.Fatal("读取doc数据失败:", err.Error())
		}
		return
	}
	content, err := os.ReadFile(s.dataPath)
	if err != nil {

		log.Fatal("读取doc数据失败:", err.Error())
	}
	var payload DataFile[T]
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
		return
	}
	s.baseList = payload.List
}

func (s *BaseService[T]) writeFile() error {
	data := &DataFile[T]{
		List: s.baseList,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal("Error during Marshal(): ", err)
		return err
	}
	file, err := os.Create(s.dataPath)
	if err != nil {
		log.Fatal("Error creating file:", err)
		return err
	}
	defer file.Close()
	_, err = file.Write(jsonData)
	if err != nil {
		log.Fatal("Error writing file:", err)
		return err
	}
	return nil
}

func (s *BaseService[T]) GetBaseList() []T {
	return s.baseList
}

func (s *BaseService[T]) Add(r T) ServiceResponse {
	r.SetId(uuid.New().String())
	r.SetCreatedAt(utils.NowStr())
	r.SetUpdatedAt(utils.NowStr())
	s.baseList = utils.Unshift(s.baseList, r)
	err := s.writeFile()
	if err != nil {
		return failure(500, err.Error(), s.baseList)
	}
	return success(s.baseList)
}
