package sensuasset

import (
	"context"
	"crypto/sha512"
	"fmt"
	"io"
	"log"
	"os"
	"path"

	"github.com/julian7/goshipdone/ctx"
	"github.com/julian7/goshipdone/modules"
)

type Asset struct {
	// AssetURL is a template for Build artifact's publicly available URL,
	// from where sensu-go agent will be able to download URLs.
	// There is no default, and this is a required parameter.
	AssetURL string `yaml:"asset_url"`
	// Build specifies which build name should be added to the asset.
	// There is no default, and this is a required parameter.
	Build string
	// ID contains the artifact's name used by later stages of the build
	// pipeline. Archives, and Publishes may refer to this name for
	// referencing build results.
	// Default: "sensu-asset".
	ID string
	// Output is where the checksum file is going to be created
	// Default: "{{.ProjectName}}-{{.Version}}-asset.yml"
	Output string
}

func Register() {
	modules.RegisterModule(&modules.ModuleRegistration{
		Stage:   "build",
		Type:    "sensu-asset",
		Factory: NewAsset,
	})
}

func NewAsset() modules.Pluggable {
	return &Asset{
		ID:     "sensu-asset",
		Output: "{{.ProjectName}}-{{.Version}}-asset.yml",
	}
}

func (mod *Asset) Run(cx context.Context) error {
	context, err := ctx.GetShipContext(cx)
	if err != nil {
		return fmt.Errorf("sensu-asset run: %w", err)
	}

	assetspec := NewAssetSpec(context.ProjectName)
	builds := context.Artifacts.ByID(mod.Build)

	td, err := modules.NewTemplate(cx)
	if err != nil {
		return fmt.Errorf("sensu-asset template: %w", err)
	}

	outfile, err := td.Parse("sensu-asset:outfile", mod.Output)
	if err != nil {
		return fmt.Errorf("parsing output: %w", err)
	}

	for _, build := range *builds {
		td.OS = build.OS
		td.Arch = build.Arch
		td.ArchiveName = build.Filename

		assetURL, err := td.Parse("sensu-asset:url", mod.AssetURL)
		if err != nil {
			return fmt.Errorf("parsing AssetURL: %w", err)
		}

		spec, err := mod.summarize(build, assetURL)
		if err != nil {
			return err
		}

		assetspec.AddBuild(spec)
	}

	location := path.Join(context.TargetDir, outfile)
	writer, err := os.Create(location)
	if err != nil {
		return fmt.Errorf("opening sensu-asset file for writing: %w", err)
	}

	defer writer.Close()

	_, err = assetspec.Write(writer)
	if err != nil {
		return fmt.Errorf("writing sensu-asset file: %w", err)
	}

	context.Artifacts.Add(&ctx.Artifact{
		ID:       mod.ID,
		Filename: outfile,
		Location: location,
	})

	log.Printf("sensu-asset file %s written.", location)
	return nil
}

func (mod *Asset) summarize(art *ctx.Artifact, url string) (*BuildSpec, error) {
	reader, err := os.Open(art.Location)
	if err != nil {
		return nil, fmt.Errorf("cannot open file for calculating SHA512 checksum: %w", err)
	}
	defer reader.Close()

	hash := sha512.New()
	if _, err := io.Copy(hash, reader); err != nil {
		return nil, fmt.Errorf("cannot read file for calculating SHA512 checksum: %w", err)
	}

	build := &BuildSpec{
		URL:    url,
		SHA512: fmt.Sprintf("%x", hash.Sum(nil)),
		Filters: []string{
			fmt.Sprintf("entity.system.os == '%s'", art.OS),
			fmt.Sprintf("entity.system.arch == '%s'", art.Arch),
		},
	}

	return build, nil
}
