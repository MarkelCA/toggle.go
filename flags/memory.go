package flags

type MemoryRepository struct {}

func NewMemoryRepository() FlagRepository {
    return MemoryRepository{}
}

var flags []Flag = make([]Flag,0)

func (r MemoryRepository) List()[]Flag {
    return flags
}

func (r MemoryRepository) Create(flag Flag) error {
    if r.Exists(flag.Name) {
        return &FlagExistsError{}
    }
    flags = append(flags,flag)
    return nil
}

func (r MemoryRepository) Exists(name string) bool {
    for _,currentFlag := range flags {
        if currentFlag.Name == name {
            return true
        }
    }
    return false
}
