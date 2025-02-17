package logic

import "os/exec"
import "fmt"
import "encoding/json"
import "runtime"
import "G19_heis2/Heis/config"

// Struct members must be public in order to be accessible by json.Marshal/.Unmarshal
// This means they must start with a capital letter, so we need to use field renaming struct tags to make them camelCase



func HallRequestAssigner(hallRequests [][2]bool, states map[string]config.HRAElevState)(map[string][][2]bool,error){

    hraExecutable := ""
    switch runtime.GOOS {
        case "linux":   hraExecutable  = "hall_request_assigner"
        case "windows": hraExecutable  = "hall_request_assigner.exe"
        default:        panic("OS not supported")
    }

    input := config.HRAInput{
        HallRequests: hallRequests,
        States: states,
            
    }

    jsonBytes, err := json.Marshal(input)
    if err != nil {
       
        return nil, fmt.Errorf("json.Marshal error: ", err)
    }
    
    ret, err := exec.Command("../hall_request_assigner/"+hraExecutable, "-i", string(jsonBytes)).CombinedOutput()
    if err != nil {
     
        return nil, fmt.Errorf("exec.Command error: %v, output: %s", err, string(ret))
    }
    
    var output map[string][][2]bool
    if err = json.Unmarshal(ret, &output)
    err != nil {
        return nil, fmt.Errorf("json.Unmarshal error: ", err)
    }
        


    return output,nil
  
}




