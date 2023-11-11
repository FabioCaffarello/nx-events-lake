package entity

type FileCatalogInterface interface {
    Save(fileCatalog *FileCatalog) error
    FindOneById(id string) (*FileCatalog, error)
    FindAll() ([]*FileCatalog, error)
    FindAllByService(service string) ([]*FileCatalog, error)
    FindAllByServiceAndSource(service string, source string) ([]*FileCatalog, error)
    DeleteOneById(id string) error
    FindOneByServiceAndSource(service string, source string) (*FileCatalog, error)
}
