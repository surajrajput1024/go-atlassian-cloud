package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/surajsinghrajput/go-atlassian-cloud/types"
)

func (c *Client) ResolveCloudID() (string, error) {
	if !empty(c.cfg.CloudID) {
		return trim(c.cfg.CloudID), nil
	}
	req, err := http.NewRequest(http.MethodGet, c.cfg.TenantInfoURL(), nil)
	if err != nil {
		return "", err
	}
	resp, err := c.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		apiErr, _ := newAPIError(resp)
		if apiErr != nil {
			return "", apiErr
		}
		return "", fmt.Errorf("tenant_info returned status %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var out types.TenantInfoResponse
	if err := json.Unmarshal(body, &out); err != nil {
		return "", err
	}
	if empty(out.CloudID) {
		return "", fmt.Errorf("tenant_info response missing cloudId")
	}
	return out.CloudID, nil
}
