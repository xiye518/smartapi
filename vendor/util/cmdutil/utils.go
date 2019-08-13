package cmdutil

import (
	"bytes"
	"errors"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/cosiner/argv"
)

// GetOutput 获取指令输出
func GetOutput(cmd string) (string, error) {
	out := bytes.NewBuffer(make([]byte, 1024))
	err := ExecCommand(cmd, out)
	if err != nil {
		return "", err
	}

	return out.String(), nil
}

// GetCmdsOutput 获取指令输出
func GetCmdsOutput(cmds [][]string) (string, error) {
	out := bytes.NewBuffer(make([]byte, 1024))
	err := PipeExecCommand(cmds, out)
	if err != nil {
		return "", err
	}

	return out.String(), nil
}

// ExecCommand 执行shell命令
func ExecCommand(cmd string, out io.Writer) error {
	args, err := argv.Argv([]rune(cmd), map[string]string{}, argv.Run)
	if err != nil {
		return err
	}

	return PipeExecCommand(args, out)
}

// GetShellOutput 获取Shell指令命令行输出
func GetShellOutput(cmd string) (string, error) {
	out := bytes.NewBuffer(make([]byte, 1024))
	err := ShellCommand(cmd, out)
	if err != nil {
		return "", err
	}

	return out.String(), nil
}

// ShellCommand 原样执行shell命令
func ShellCommand(cmd string, out io.Writer) error {
	in := strings.NewReader(cmd)
	c := exec.Command("sh")
	c.Stdin = in
	c.Stdout = out

	if err := c.Start(); err != nil {
		return err
	}
	return c.Wait()
}

// PipeExecCommand 执行管道命令
func PipeExecCommand(cmds [][]string, out io.Writer) error {
	if len(cmds) < 1 {
		return errors.New("PipeExecCommand get no command")
	}

	var commands []*exec.Cmd
	for _, cmd := range cmds {
		if len(cmd) < 1 {
			return errors.New("PipeExecCommand get null command")
		}
		commands = append(commands, exec.Command(cmd[0], cmd[1:]...))
	}

	// 连接命令
	for i := 1; i < len(commands); i++ {
		commands[i].Stdin, _ = commands[i-1].StdoutPipe()

	}
	commands[len(commands)-1].Stdout = out

	// 开启除第一个命令外的其他命令
	for i := 1; i < len(commands); i++ {
		err := commands[i].Start()
		if err != nil {
			return err
		}
	}

	// 开启第一个命令
	if err := commands[0].Start(); err != nil {
		return err
	}

	// 等待其他命令结束
	for i := 0; i < len(commands); i++ {
		if err := commands[i].Wait(); err != nil {
			return err
		}
	}

	return nil
}

// IsExists 判断路径是否存在
func IsExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}
