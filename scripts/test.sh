#!/bin/bash

set -Eeuxo pipefail

cd "$( dirname "${BASH_SOURCE[0]}" )/.."

source scripts/prep.sh

rm -rf config/client/*.proc.yml

for f in config/client/*
do
  echo "Processing $f..."
  envsubst < "${f}" > "${f}.proc.yml"
done

typescript () {
  echo "Testing TypeScript..."

  dir="clients/${PROJECT}/typescript"
  (cd "$dir"; npm i; npm run build)
}

java () {
  echo "Testing Java..."

  dir="clients/${PROJECT}/java"
  (cd "$dir"; mvn test-compile)
}

php() {
  echo "Testing PHP..."

  dir="clients/${PROJECT}/php"
  (cd "$dir"; composer install; ./vendor/bin/phpunit)
}

python () {
  echo "Generating Python..."

  dir="clients/${PROJECT}/python"
  (cd "$dir"; pip install -r requirements.txt; pip install -r test-requirements.txt; pytest --cov="$PYTHON_PACKAGE_NAME")
}

ruby () {
  echo "Generating Ruby..."

  dir="clients/${PROJECT}/ruby"
  (cd "$dir"; rm "${RUBY_PROJECT_NAME}-${GEM_VERSION}.gem" || true; bundle install --path vendor/bundle; bundle exec rspec; gem build "${RUBY_PROJECT_NAME}.gemspec"; gem install "${RUBY_PROJECT_NAME}-${GEM_VERSION}.gem")
}

golang () {
  echo "Generating Golang..."

  dir="clients/${PROJECT}/go"
  (cd "${dir}"; go mod tidy; go build -o "$(mktemp)" .)
}

dotnet () {
  echo "Generating dotnet..."

  dir="clients/${PROJECT}/dotnet"
  (cd "$dir"; dotnet build -c Release; dotnet test -c Release)
}

dart () {
  echo "Generating Dart..."

  dir="clients/${PROJECT}/dart"
  (cd "$dir"; dart test .)
}

rust () {
  echo "Generating Rust..."

  dir="clients/${PROJECT}/rust"
  (cd "$dir"; cargo test)
}

typescript
rust
golang
java
php
python
ruby
dotnet
dart