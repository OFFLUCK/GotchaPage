package tests

import (
	"fmt"
	"gotchaPage/requesters"
	"testing"
)

func TestNewSocialNetworkRequester(t *testing.T) {
	testCases := []struct {
		name     string
		mainURL  string
		nickname string
		expected *requesters.SocialNetworkRequester
	}{
		{
			name:     "VK",
			mainURL:  "vk.com",
			nickname: "olegsama",
			expected: &requesters.SocialNetworkRequester{
				Name:     "VK",
				MainURL:  "vk.com",
				Nickname: "olegsama",
			},
		},
		{
			name:     "Github",
			mainURL:  "github.com",
			nickname: "OFFLUCK",
			expected: &requesters.SocialNetworkRequester{
				Name:     "Github",
				MainURL:  "github.com",
				Nickname: "OFFLUCK",
			},
		},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprintf("Test№%d", index), func(t *testing.T) {
			got := requesters.NewSocialNetworkRequester(
				testCase.name,
				testCase.mainURL,
				testCase.nickname,
			)

			if *testCase.expected != *got {
				t.Errorf("Expected: %s; got: %s\n", testCase.expected, got)
			}
		})
	}
}

func TestGetName(t *testing.T) {
	testCases := []struct {
		requester *requesters.SocialNetworkRequester
		expected  string
	}{
		{
			requesters.NewSocialNetworkRequester(
				"VK",
				"vk.com",
				"olegsama",
			),
			"VK",
		},
		{
			requesters.NewSocialNetworkRequester(
				"Github",
				"github.com",
				"OFFLUCK",
			),
			"Github",
		},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprintf("Test№%d", index), func(t *testing.T) {
			got := testCase.requester.GetName()
			if testCase.expected != got {
				t.Errorf("Expected: %s; got: %s\n", testCase.expected, got)
			}
		})
	}
}

func TestGetInfo(t *testing.T) {
	testCases := []struct {
		requester    *requesters.SocialNetworkRequester
		expectedLink string
		expectedName string
	}{
		{
			requesters.NewSocialNetworkRequester(
				"VK",
				"vk.com",
				"olegsama",
			),
			"https://vk.com/olegsama",
			"Олег Сидоренков | ВКонтакте",
		},
		{
			requesters.NewSocialNetworkRequester(
				"Github",
				"github.com",
				"OFFLUCK",
			),
			"https://github.com/OFFLUCK",
			"OFFLUCK (Oleg) · GitHub",
		},
		{
			requesters.NewSocialNetworkRequester(
				"VK",
				"vk.com",
				"dmvdfcjdjk123211hj23123bhwhb1hb3j",
			),
			"page not found",
			"",
		},
		{
			requesters.NewSocialNetworkRequester(
				"Github",
				"github.com",
				"dm5vdfcj31djk2321151e34123214211hj2323e123sd211342bhwhb1hb3j",
			),
			"page not found",
			"",
		},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprintf("Test№%d", index), func(t *testing.T) {
			gotLink, gotName, err := testCase.requester.GetInfo()
			if err == nil {
				if gotLink != testCase.expectedLink || gotName != testCase.expectedName {
					t.Errorf("Expected: %s, %s; got: %s, %s\n", testCase.expectedLink, testCase.expectedName, gotLink, gotName)
				}
				return
			}

			if err.Error() == "page not found" {
				if gotLink != testCase.expectedLink {
					t.Errorf("Expected: %s, %s; got: %s, %s\n", testCase.expectedLink, testCase.expectedName, gotLink, gotName)
				}
				return
			}

			t.Fatalf("Unexpected error: %s", err)
		})
	}
}
