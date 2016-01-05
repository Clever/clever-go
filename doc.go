// Package clever is a Go library for the Clever API: https://clever.com/developers/docs.
//
//
// Usage
//
//			package main
//
//			import (
//				"golang.org/x/oauth2"
//				clevergo "gopkg.in/Clever/clever-go.v1"
//				"log"
//			)
//
//			func main() {
//				t := &oauth.Transport{
//					Source: oauth.StaticTokenSource(&oauth2.Token{AccessToken: "DEMO_TOKEN"}),
//				}
//				client := &http.client{Transport: t}
//				clever := clevergo.New(client, "https://api.clever.com")
//				paged := clever.QueryAll("/v1.1/districts", nil)
//				for paged.Next() {
//					var district clevergo.District
//					if err := paged.Scan(&district); err != nil {
//						log.Fatal(err)
//					}
//					log.Println(district)
//				}
//				if err := paged.Error(); err != nil {
//					log.Fatal(err)
//				}
//			}
//
package clever
