# csvfilter

[![Build Status](https://secure.travis-ci.org/jehiah/csvfilter.png?branch=master)](http://travis-ci.org/jehiah/csvfilter) [![GitHub release](https://img.shields.io/github/release/jehiah/csvfilter.svg)](https://github.com/jehiah/csvfilter/releases/latest)

`csvfilter` provides an easy way to extract specific CSV columns in a unix pipeline

To extract just the 0th and 2nd CSV column (skipping the 1st column) you could run this:

```bash
cat data.csv | csvfilter -c 0,2
```