# National Drug Code Directory

This package can be used to download and parse the NDCD database.
I am not affiliated with the FDA in any way. I initially wrote this to practice
writing text parsers; however, I have used it in conjunction with the openfda api
for data mining.

Feel free to submit pull requests.

## About NDCD
[fda.gov/Drugs/InformationOnDrugs](http://www.fda.gov/Drugs/InformationOnDrugs/ucm142438.htm)

## TODO
Write methods
  1. Refactor panics to instead return an error value
  2. Support parsing package.txt into a ProductPackage data structure linked to the proper product
