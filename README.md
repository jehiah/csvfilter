# csvfilter

[![Go package](https://github.com/jehiah/csvfilter/actions/workflows/ci.yml/badge.svg)](https://github.com/jehiah/csvfilter/actions/workflows/ci.yml) [![GitHub release](https://img.shields.io/github/release/jehiah/csvfilter.svg)](https://github.com/jehiah/csvfilter/releases/latest)

## Install

```bash
go install github.com/jehiah/csvfilter@latest
```

Pre-built binaries are available on the [releases page](https://github.com/jehiah/csvfilter/releases/latest).

## Usage

`csvfilter` provides an easy way to extract specific CSV columns in a unix pipeline

To extract just the 0th and 2nd CSV column (skipping the 1st column) you could run this:

```bash
cat data.csv | csvfilter -c 0,2
```