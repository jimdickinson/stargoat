package stargoat

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"time"

	"encoding/json"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/jimdickinson/stargoat/internal/generated/client"
	"github.com/jimdickinson/stargoat/internal/generated/client/documents"
)

type Client struct {
	Ctx               context.Context
	Token             StargateToken
	StayAuthenticated bool
	impl              *client.Stargate
	httpClient        *http.Client
	shutdown          bool
}

func NewClient(sgURL string, token StargateToken, stayAuthenticated bool, httpClient *http.Client) (*Client, error) {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	c := Client{
		Ctx:               context.Background(),
		Token:             token,
		StayAuthenticated: stayAuthenticated,
		shutdown:          false,
	}
	u, err := url.Parse(sgURL)
	if err != nil {
		return nil, err
	}
	transport := httptransport.NewWithClient(u.Host, u.Path, []string{u.Scheme}, httpClient)
	transport.EnableConnectionReuse()
	c.impl = client.New(transport, nil)
	c.httpClient = httpClient
	if stayAuthenticated {
		go c.tokenKeepAlive()
	}
	return &c, nil
}

const refreshInterval = 5 * time.Minute
const tickerInterval = 1 * time.Minute

func (c *Client) tokenKeepAlive() {
	lastSuccess := time.Now()
	ticker := time.NewTicker(tickerInterval)
	defer ticker.Stop()
	for {
		_ = <-ticker.C
		if c.shutdown {
			break
		}
		if time.Until(lastSuccess.Add(refreshInterval)) < 0 {
			params := &documents.GetAllNamespacesParams{
				XCassandraToken: c.Token.AuthToken,
				Context:         c.Ctx,
			}
			resp, err := c.impl.Documents.GetAllNamespaces(params)
			if err != nil {
				log.Printf("%+v", err)
			} else if resp != nil {
				lastSuccess = time.Now()
				log.Printf("%+v", resp)
			} else {
				log.Printf("both nil")
			}
		}
	}
}

func (c *Client) PostDoc(namespace, collection string, document interface{}) (string, error) {
	b, err := json.Marshal(document)
	if err != nil {
		return "", err
	}
	params := &documents.PostDocParams{}
	params.Context = c.Ctx
	params.HTTPClient = c.httpClient
	params.XCassandraToken = c.Token.AuthToken
	params.NamespaceID = namespace
	params.CollectionID = collection
	params.Body = json.RawMessage(b)
	response, err := c.impl.Documents.PostDoc(params)
	if err != nil {
		return "", err
	}
	return response.Payload.DocumentID, nil
}

func (c *Client) PutDoc(namespace, collection, id string, document interface{}) (string, error) {
	b, err := json.Marshal(document)
	if err != nil {
		return "", err
	}
	params := &documents.PutDocParams{}
	params.Context = c.Ctx
	params.HTTPClient = c.httpClient
	params.XCassandraToken = c.Token.AuthToken
	params.NamespaceID = namespace
	params.CollectionID = collection
	params.DocumentID = id
	params.Body = json.RawMessage(b)
	response, err := c.impl.Documents.PutDoc(params)
	if err != nil {
		return "", err
	}
	return response.Payload.DocumentID, nil
}

func (c *Client) PatchDoc(namespace, collection, id string, document interface{}) (string, error) {
	b, err := json.Marshal(document)
	if err != nil {
		return "", err
	}
	params := &documents.PatchDocParams{}
	params.Context = c.Ctx
	params.HTTPClient = c.httpClient
	params.XCassandraToken = c.Token.AuthToken
	params.NamespaceID = namespace
	params.CollectionID = collection
	params.DocumentID = id
	params.Body = json.RawMessage(b)
	response, err := c.impl.Documents.PatchDoc(params)
	if err != nil {
		return "", err
	}
	return response.Payload.DocumentID, nil
}

func (c *Client) GetDoc(namespace, collection, id string) (interface{}, error) {
	params := &documents.GetDocParams{}
	params.Context = c.Ctx
	params.HTTPClient = c.httpClient
	params.XCassandraToken = c.Token.AuthToken
	params.NamespaceID = namespace
	params.CollectionID = collection
	params.DocumentID = id

	docOK, docNoContent, err := c.impl.Documents.GetDoc(params)
	if err != nil {
		return nil, err
	}
	if docNoContent != nil {
		return nil, nil
	}
	return docOK.Payload.Data, nil
}

func (c *Client) DeleteDoc(namespace, collection, id string) (bool, error) {
	params := &documents.DeleteDocParams{}
	params.Context = c.Ctx
	params.HTTPClient = c.httpClient
	params.XCassandraToken = c.Token.AuthToken
	params.NamespaceID = namespace
	params.CollectionID = collection
	params.DocumentID = id

	docNoContent, err := c.impl.Documents.DeleteDoc(params)
	if err != nil {
		return false, err
	}
	return docNoContent != nil, nil
}

// func (c *Client) CreateCollection(name string) error {
// 	params := documents.CreateCollectionParams{}
// 	params.XCassandraToken = c.Token.AuthToken
// 	params.Body = name
// 	_, err := c.impl.Documents.CreateCollection(&params)
// 	return err
// }

// func (c *Client) CreateNamespace(params *CreateNamespaceParams) (*CreateNamespaceCreated, error)    {}
// func (c *Client) DeleteCollection(params *DeleteCollectionParams) (*DeleteCollectionNoContent, error) {}
// func (c *Client) DeleteDocPath(params *DeleteDocPathParams) (*DeleteDocPathNoContent, error)       {}
// func (c *Client) DeleteNamespace(params *DeleteNamespaceParams) (*DeleteNamespaceNoContent, error) {}
// func (c *Client) GetAllNamespaces(params *GetAllNamespacesParams) (*GetAllNamespacesOK, error)     {}
// func (c *Client) GetCollections(params *GetCollectionsParams) (*GetCollectionsOK, error)           {}
// func (c *Client) GetDoc(params *GetDocParams) (*GetDocOK, *GetDocNoContent, error)                 {}
// func (c *Client) GetDocPath(params *GetDocPathParams) (*GetDocPathOK, *GetDocPathNoContent, error) {}
// func (c *Client) GetOneNamespace(params *GetOneNamespaceParams) (*GetOneNamespaceOK, error)        {}
// func (c *Client) PatchDocPath(params *PatchDocPathParams) (*PatchDocPathOK, error)                 {}
// func (c *Client) PutDocPath(params *PutDocPathParams) (*PutDocPathOK, error)                       {}
// func (c *Client) SearchDoc(params *SearchDocParams) (*SearchDocOK, *SearchDocNoContent, error)     {}
// func (c *Client) UpgradeCollection(params *UpgradeCollectionParams) (*UpgradeCollectionOK, error)  {}
// func (c *Client) SetTransport(transport runtime.ClientTransport)                                   {}
