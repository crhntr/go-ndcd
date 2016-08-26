#National Drug Code Directory

This package can be used to download and parse the NDCD database.
I am not affiliated with the FDA in any way. I initially wrote this to practice
writing text parsers.

Feel free to submit pull requests.

##About NDCD###
[fda.gov/Drugs/InformationOnDrugs](http://www.fda.gov/Drugs/InformationOnDrugs/ucm142438.htm)

##TODO##
Write methods
  1. func (product Product) Valid() bool  
  2. func (labeler Labeler) Valid() bool
  
  3. Refactor tests to be more modular.
  4. Write tests that check for panics.
  4. Refactor panics to instead return an error value
