import json
import os
from functools import cmp_to_key
from operator import itemgetter

def sort_by_repetitions(element1, element2):
  """
  Custom key comparison function for sorting by repetitions (descending order).
  """
  return (element2[1] > element1[1]) - (element2[1] < element1[1])

def main():
  # Replace "data.json" with the actual path to your JSON file
  input_filename = "src/ETL/explorer/match_counts.json"
  # Specify the output filename (avoid overwriting the original file)
  output_filename = "src/ETL/explorer/sorted_data.json"

  # Check if input file exists
  if not os.path.exists(input_filename):
    print(f"Error: File '{input_filename}' does not exist.")
    return

  # Read JSON file
  try:
    with open(input_filename, "r") as json_file:
      data = json.load(json_file)
  except FileNotFoundError:
    print(f"Error: Could not open file '{input_filename}'.")
    return
  except json.JSONDecodeError as e:
    print(f"Error: Invalid JSON format in file '{input_filename}'.")
    print(f"Details: {e}")
    return

  # Sort the data by repetitions (descending order)
  sorted_data = dict(sorted(data.items(), key=cmp_to_key(sort_by_repetitions)))

  # Save the sorted data to a new file
  try:
    with open(output_filename, "w") as json_file:
      json.dump(sorted_data, json_file, indent=4)
  except OSError as e:
    print(f"Error: Could not write to file '{output_filename}'.")
    print(f"Details: {e}")
    return

  print(f"Successfully saved sorted data to '{output_filename}'.")

if __name__ == "__main__":
  main()
