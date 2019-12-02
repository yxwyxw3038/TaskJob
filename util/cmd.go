package util
import (
    "os/exec"
)
func  Run(path string, mid string, id string) error {
	// arg := []string{mid, id}
	// cmd := exec.Command("cmd.exe", "/c", "start "+path)
	// context := mid +" "+id+ " >> " + path
	// cmd := exec.Command("cmd.exe", `/c`+context)
	context := "start "+path+" "+mid +" "+id
	cmd := exec.Command("cmd.exe","/c", context)
	if err := cmd.Run(); err != nil {
       return err
	}  
	// cmd = exec.Command("sleep", "10")
	// if err1 := cmd.Run(); err1 != nil {
	// 	return err1
	//  }  
	return nil
}