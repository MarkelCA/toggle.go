package storage
//
// // REFACTOR WITHOUT THE flags dependency
//
// // import (
// // 	"time"
// //     "reflect"
// // 	"github.com/markelca/toggles/flags"
// // )
// //
// type MemoryRepository struct {}
// //
// func NewMemoryRepository() CacheClient {
//     return MemoryRepository{}
// }
// //
// var flagsStorage map[string]interface{} = make(map[string]interface{})
// // var flagsStorage []interface{} = make([]interface{},0)
// // var flagsStorage []flags.Flag = make([]flags.Flag,0)
// //
// func(r MemoryRepository) Keys() ([]string, error) {
//     result := make([]string, 0)
//     for i,_ := range flagsStorage {
//         result = append(result, i)
//     }
//     return result,nil
// }
// //
// func(r MemoryRepository) Get(key string) (interface{}, error) {
//     for k,v   := range flagsStorage {
//         if k ==  key {
//             return v,nil
//         }
//     }
//     return false,nil
// }
//
// //
// // func (r MemoryRepository) Set(key string, value interface{}, expiration time.Duration) error {
// //     v := reflect.ValueOf(value).Bool()
// //     for i,currentFlag := range flagsStorage {
// //         if currentFlag.Name == key {
// //             flagsStorage[i].Value = v
// //             return nil
// //         }
// //     }
// //     // If it doesn't find it it adds it
// //     flagsStorage = append(flagsStorage,flags.Flag{Name:key,Value:v,})
// //     return nil
// // }
// //
// // func (r MemoryRepository) Exists(name string) (bool,error) {
// //     for _,currentFlag := range flagsStorage {
// //         if currentFlag.Name == name {
// //             return true,nil
// //         }
// //     }
// //     return false,nil
// // }
// //
