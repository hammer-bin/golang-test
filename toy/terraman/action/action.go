package action

import (
	"fmt"
	"os/exec"
)

const kubectl = "kubectl"
const baseDir = "/tmp/terraform/"

func GetTodo() {
	/*
	   1. kubectl cp로 파일 가져오기
	      kubectl cp -n mariadb paas-ta-container-platform-mariadb-0:var/lib/apt/extended_states /home/ubuntu/aa   --파일 복사
	      kubectl cp -n mariadb paas-ta-container-platform-mariadb-0:bitnami/ /home/ubuntu/bitnami    --폴더 복사
	   2. 파일 디렉토리 조정
	   3. IaaS에 따라 provider.tf 파일 정의
	   4. terraform init 실행
	   5. terraform plan
	   6. terraform apply
	   7. Infra생성 후 생성된 Instance IP 알아오기
	   8. kubespray_var.sh 파일 작성하기
	   9. source deploy_kubespray.sh 실행
	   10. 생성 상태 전송 --> API
	   11. 완료 후 종료
	*/
	fmt.Println("get To do...!!")

	result, err := execCommandOutput(kubectl, "cp", "-n", "mariadb", "paas-ta-container-platform-mariadb-0:bitnami/", baseDir+"/bitnami2")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("파일 복사가 완료되었습니다.", result)
	}
}

func execCommandOutput(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)

	//cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := cmd.CombinedOutput()
	return string(output), err
}
