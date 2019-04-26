OpenPing
===

OpenPing is a monitor for websites. It has 2 basic functions:

- Ensure a URL returns a 200
- Ensure a URL doesn't exceed a set deviation.

The components of OpenPing are:

- A binary that scrapes its endpoint targets
- A document store (presently `mongodb`) to store the last n iterations of the document
