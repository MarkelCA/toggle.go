package flags

type MemoryRepository struct {}

func NewMemoryRepository() FlagRepository {
    return MemoryRepository{}
}

var flags []Flag = make([]Flag,0)

func (r MemoryRepository) List()([]Flag, error) {
    return flags,nil
}

func (r MemoryRepository) Create(flag Flag) error {
    result,err := r.Exists(flag.Name)

    if err != nil{
        return err
    }
    if result {
        return &FlagExistsError{}
    }
    flags = append(flags,flag)
    return nil
}

func (r MemoryRepository) Exists(name string) (bool,error) {
    for _,currentFlag := range flags {
        if currentFlag.Name == name {
            return true,nil
        }
    }
    return false,nil
}

func (r MemoryRepository) Update(name string, value bool) error {
    for i := range flags {
        if flags[i].Name == name {
            flags[i].Value = value
            break
        }
    }
    return nil
}
