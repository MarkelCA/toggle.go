package flags


type FlagService struct {
    Repository FlagRepository
}

func NewFlagService(r FlagRepository) FlagService {
    return FlagService{r}
}

func (s FlagService) Get(key string) (bool,error) {
    return s.Repository.Get(key)
}

func (s FlagService) Create(f Flag) error {
    exists,err := s.Exists(f.Name)
    if err != nil {
        return err
    } else if exists {
        return FlagAlreadyExistsError
    } 

    err = s.Repository.Set(f,0)
    if err != nil {
        return err
    }

    return nil
}

func (s FlagService) Update(name string, value bool) error {
    exists,err := s.Exists(name)
    if err != nil {
        return err
    } else if !exists {
        return FlagNotFoundError
    }
    err = s.Repository.Set(Flag{name,value},0)
    if err != nil {
        return err
    }
    return nil
}

func (s FlagService) Exists(key string) (bool,error) {
    return s.Repository.Exists(key)
}
 
func (s FlagService) List()([]Flag, error) {
    keys,err := s.Repository.Keys()
    if err != nil {
        return nil,err
    }

    result := make([]Flag,len(keys))

    for i,key := range keys {
        val,err := s.Repository.Get(key)
        if err != nil {
            return nil,err
        }
        result[i] = Flag{
            Name: key,
            Value: val,
        }
    }
    return result,nil
}

