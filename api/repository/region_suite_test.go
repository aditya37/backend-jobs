/*
 * File Name repository_suite_test.go
 * Created on Sat Sep 12 2020
 *
 * Copyright (c) 2020
 */
package repository_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRepository(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Repository Suite")
}
