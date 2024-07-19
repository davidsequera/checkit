import os
import json
from collections import defaultdict

def count_files_recursively(folder_path):
    total_files = 0
    for root, dirs, files in os.walk(folder_path):
        total_files += len(files)
    return total_files

def create_sorted_folder_json(directory_path, output_file):
    folder_dict = defaultdict(int)

    # Count files in each first-level folder
    for item in os.listdir(directory_path):
        item_path = os.path.join(directory_path, item)
        if os.path.isdir(item_path):
            file_count = count_files_recursively(item_path)
            folder_dict[item] = file_count

    # Sort folders by file count (descending order)
    sorted_folders = sorted(folder_dict.items(), key=lambda x: x[1], reverse=True)

    # Create a list of dictionaries for JSON output
    folder_list = [{"folder": folder, "file_count": count} for folder, count in sorted_folders]

    # Write to JSON file
    with open(output_file, 'w') as f:
        json.dump(folder_list, f, indent=2)

    print(f"JSON file '{output_file}' has been created successfully.")

# Example usage
directory_path = "db/maildir"
output_file = "folder_file_counts.json"
create_sorted_folder_json(directory_path, output_file)