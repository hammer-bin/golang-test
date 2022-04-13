package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
)

const baseDir = "/tmp/terraform/"
const kubectl = "kubectl"

var iaas = map[string]*IaaS{}

type IaaS struct {
	csp string `json:"csp"`
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, r)
	})
}

func execCommandOutput(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)

	//cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := cmd.CombinedOutput()
	return string(output), err
}

func getTodo() {
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

	result, err := execCommandOutput(kubectl, "cp", "-n", "mariadb", "paas-ta-container-platform-mariadb-0:bitnami/", "/home/ubuntu/bitnami2")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("파일 복사가 완료되었습니다.", result)
	}
}

func rest() {
	mux := http.NewServeMux()

	userHandler := http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			json.NewEncoder(wr).Encode(iaas)
			fmt.Println("user : ", iaas)
			getTodo() // 신호를 받았을 때 액션 처리
		case http.MethodPost:
			var infra IaaS
			json.NewDecoder(r.Body).Decode(&infra)
			iaas[infra.csp] = &infra
			json.NewEncoder(wr).Encode(infra)
		}
	})

	mux.Handle("/users", jsonContentTypeMiddleware(userHandler))
	http.ListenAndServe(":8080", mux)
}

func main() {

	rest() // API 서버 동작

}
