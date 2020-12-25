package repository_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestEmploye(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Employe Suite")
}
