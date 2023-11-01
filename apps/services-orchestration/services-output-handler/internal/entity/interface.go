package entity

type ServiceOutputInterface interface {
     Save(serviceOutput *ServiceOutput, service string) error
     FindOneByIdAndService(id string, service string) (*ServiceOutput, error)
     FindAllByService(service string) ([]*ServiceOutput, error)
     FindAllByServiceAndSource(service string, source string) ([]*ServiceOutput, error)
     FindAllByServiceAndSourceAndContext(service string, source string, contextEnv string) ([]*ServiceOutput, error)
}
