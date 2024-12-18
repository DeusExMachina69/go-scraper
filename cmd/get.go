package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

const (
    // schemas
    httpSchema  = "http://"
    httpsSchema = "https://"

    // TLD (.com, .org, etc.)
    comTld      = ".com"
    orgTld      = ".org"
    netTld      = ".net"
    eduTld      = ".edu"
    govTld      = ".gov"
)

var (
    schemas []string
    tlds    []string
)

func init() {
    schemas = append(schemas, httpSchema, httpsSchema)
    tlds = append(tlds, comTld, orgTld, netTld, eduTld, govTld)
}

type Links struct {
    Links   []string `json:"links"`
}

type Response struct {
    doc     *html.Node
    Links
}

func Get(url string) (*Response, error) {
    res, err := http.Get(url) 
    if err != nil {
        res.Body.Close()
        return nil, err
    }
    defer res.Body.Close()
    d, err := html.Parse(res.Body)
    if err != nil {
        return nil, err
    }
    return &Response{
        doc:    d,
        Links:  parseDoc(d),
    }, nil
}

func (r *Response) Json() error {
    js, err := json.MarshalIndent(r.Links, "", "    ")     
    if err != nil {
        fmt.Printf("Error marshalling json: %s", err)
        return err
    }
    fmt.Println(string(append(js, '\n')))
    return nil
}

func parseDoc(doc *html.Node) Links {
    l := Links{}
    for n := range doc.Descendants() {
        if n.Type == html.ElementNode && n.DataAtom == atom.A {
            for _, val := range n.Attr {
                if val.Key == "href" {
                    if hasSchema(val.Val) {
                        l.Links = append(l.Links, val.Val)
                    }
                } 
            }
        }
    }
    return l
}

func (r *Response) Print() {
    for _, v := range r.Links.Links {
        fmt.Println(v)
    }
}

func hasSchema(url string) bool {
    url = strings.Trim(url, " ")
    for _, s := range schemas {
        if strings.HasPrefix(url, s) {
            return true
        }
    }
    return false
}

func prependSchema(url string) string {
    return strings.Trim(httpsSchema + url, " ")
}

func hasTLD(url string) bool {
    url = strings.Trim(url, " ")
    for _, t := range tlds {
        if strings.Contains(url, t) {
            return true
        }
    }
    return false
}

func appendTLD(url string) string {
    return strings.Trim(url + comTld, " ")
}
