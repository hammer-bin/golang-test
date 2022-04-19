provider "google" {
	credentials = "${file("paasta-cp-3728429.json")}"
	project = "paasta-cp"
	region = "northamerica-northeast2"
}
