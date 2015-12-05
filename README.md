# tidyblob

tidyblob examines a BOSH release directory and reports any 'stale' blobs - blobs that are listed in the `config/blobs.yml` file but that are not referenced in any package's `spec` file.
