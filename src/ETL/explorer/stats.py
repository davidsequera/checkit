import json
import statistics

def read_json(filename):
    with open(filename, 'r') as f:
        data = json.load(f)
    return data

def calculate_statistics(data):
    file_counts = [entry["file_count"] for entry in data]
    
    # Calculate mean, standard deviation, and other statistics
    mean = statistics.mean(file_counts)
    std_dev = statistics.stdev(file_counts)
    var = statistics.variance(file_counts)
    min_value = min(file_counts)
    max_value = max(file_counts)
    range_value = max_value - min_value
    
    # Calculate normalized standard deviation
    if mean != 0:
        normalized_std_dev = std_dev / mean
    else:
        normalized_std_dev = None
    
    return {
        "mean": mean,
        "standard_deviation": std_dev,
        "variance": var,
        "min_value": min_value,
        "max_value": max_value,
        "range": range_value,
        "normalized_standard_deviation": normalized_std_dev
    }

if __name__ == "__main__":
    filename = "./src/etl/explorer/result/folder_file_counts.json"  # Replace with your actual file name
    
    # Read data from JSON file
    data = read_json(filename)
    
    # Calculate statistics
    stats = calculate_statistics(data)
    
    # Print the results
    print(f"Mean file count: {stats['mean']}")
    print(f"Standard deviation of file count: {stats['standard_deviation']}")
    print(f"Variance of file count: {stats['variance']}")
    print(f"Min file count: {stats['min_value']}")
    print(f"Max file count: {stats['max_value']}")
    print(f"Range of file count: {stats['range']}")
    if stats['normalized_standard_deviation'] is not None:
        print(f"Normalized standard deviation: {stats['normalized_standard_deviation']:.4f}")
    else:
        print("Cannot calculate normalized standard deviation (mean is zero)")
