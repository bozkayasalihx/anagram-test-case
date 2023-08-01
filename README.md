# AnagramFinder

## Getting Started

### Installing

To install the AnagramFinder application, follow these steps:

1. Clone the repository:

   ```bash
   git clone https://github.com/bozkayasalihx/anagram-test-case
   cd anagram-test-case
   ```

2. build anagram-finder
   ```bash
   go build -o anagram-finder
   ```

3. run anagram-finder
   ```bash
   ./anagram-finder your_input_file
   ```

## Example
   Create an `.txt` file with contains such as following contents
   ```text
      abalone
      abalones
      abandon
      abandoned
      abandonedly
      abandonee
      abandoner
      abandoners
      abandoning
      abandonment
      abandonments
      abandons
   ```

## Running in Kubernetes
 To deploy the anagram-finder application in a kubernetes environment
 
 1. apply `kubernetes/setup` manifests for pre-setup
 
 2. apply `kubernetes` to deploy anagram-finder job
 
