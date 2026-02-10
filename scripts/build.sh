#!/usr/bin/env bash
set -euo pipefail

VERSION="${VERSION:-dev}"
COMMIT_SHA="${COMMIT_SHA:-$(git rev-parse --short HEAD 2>/dev/null || echo "unknown")}"
BUILD_DATE="${BUILD_DATE:-$(date -u +"%Y-%m-%dT%H:%M:%SZ")}"

LDFLAGS="-X main.Version=${VERSION} -X main.CommitSHA=${COMMIT_SHA} -X main.BuildDate=${BUILD_DATE}"

mkdir -p bin

PLATFORMS=(
  "darwin/arm64"
  "darwin/amd64"
  "linux/amd64"
  "linux/arm64"
  "windows/amd64"
)

for platform in "${PLATFORMS[@]}"; do
  GOOS="${platform%/*}"
  GOARCH="${platform#*/}"
  OUTPUT="bin/rooda-${GOOS}-${GOARCH}"
  
  if [ "$GOOS" = "windows" ]; then
    OUTPUT="${OUTPUT}.exe"
  fi
  
  echo "Building for ${GOOS}/${GOARCH}..."
  GOOS="$GOOS" GOARCH="$GOARCH" go build -ldflags "${LDFLAGS}" -o "$OUTPUT" ./cmd/rooda
done

cd bin
if command -v sha256sum >/dev/null 2>&1; then
  sha256sum rooda-* > checksums.txt
else
  shasum -a 256 rooda-* > checksums.txt
fi
cd ..

echo "Built rooda ${VERSION} (${COMMIT_SHA}) at ${BUILD_DATE} for ${#PLATFORMS[@]} platforms"
echo "Checksums generated in bin/checksums.txt"
