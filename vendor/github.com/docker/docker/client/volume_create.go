package client // import "github.com/docker/docker/client"

import (
	"encoding/json"

	"github.com/docker/docker/api/types"
	volumetypes "github.com/docker/docker/api/types/volume"
	"golang.org/x/net/context"
)

// VolumeCreate creates a volume in the docker host.
func (cli *Client) VolumeCreate(ctx context.Context, options volumetypes.VolumesCreateBody) (types.Volume, error) {
	var volume types.Volume
	resp, err := cli.post(ctx, "/volumes/create", nil, options, nil)
	if err != nil {
		return volume, err
	}
	err = json.NewDecoder(resp.body).Decode(&volume)
	ensureReaderClosed(resp)
	return volume, err
}
