import csv
import sys


def split_data(input_data):
    result = []
    i = 0
    while i < len(input_data):
        # Check for '0x' pattern
        if i + 2 <= len(input_data) and input_data[i] == '0' and input_data[i+1].isalnum():
            result.append(' ' + input_data[i:i+2] + ' ')
            i += 2
        # Check for 'x0' pattern
        elif i + 2 <= len(input_data) and input_data[i+1] == '0' and input_data[i].isalnum():
            result.append(' ' + input_data[i:i+2] + ' ')
            i += 2
        # Check for '00' pattern
        elif i + 2 <= len(input_data) and input_data[i:i+2] == '00':
            result.append(' ' + input_data[i:i+2] + ' ')
            i += 2
        else:
            result.append(input_data[i:i+2])
            i += 2

    return ''.join(result).strip()

def remove_extra_spaces(input_string):
    # Split the string by spaces and then join them back with a single space
    return ' '.join(input_string.split())

def remove_00and0x_spaces(formatted_string):
    parts = formatted_string.split()
    # Processing the sequence to remove spaces as per the condition
    processed_parts = []
    for i in range(len(parts)):
        # Add the current part
        processed_parts.append(parts[i])
        # If the current part is '00' and there is a next part which is not '00', skip adding a space after '00'
        if parts[i] == '00' and i + 1 < len(parts) and parts[i + 1] != '00':
            continue
        else:
            # Add a space after the current part if it's not the condition above
            processed_parts.append(' ')

    processed_sequence = ''.join(processed_parts).strip()  # strip to remove the trailing space if any
    return processed_sequence

def write_multiple_lines_to_out_file(filename, data_lines):
    with open(filename, 'w') as f:
        f.write('Hexstream,Split Indexes,Splited Hexstream\n')
        for hexstream in data_lines:
            # Calculate split indexes based on spaces (ignoring spaces)
            split_indexes = []
            index = 0
            split_indexes.append(str(index))
            for char in hexstream:
                if char == ' ':
                    split_indexes.append(str(index//2))
                if char != ' ':  # Only increment index if not a space
                    index += 1

            # Format the split indexes and splited hexstream
            splited_hexstream = hexstream.replace(' ', '')
            second_line = f'{splited_hexstream},"[{", ".join(split_indexes)}]",{hexstream}\n'
            f.write(second_line)


# Read input data from a .out file
input_filename = '/root/message_split_analysis/tshark/hexstream/dnp3.out'
output_filename = '/root/message_split_analysis/tshark/zwl_out/dnp3.out'

with open(input_filename, 'r') as file:
    data_lines = file.readlines()

# Process each line from the input file
processed_lines = []
for line in data_lines:
    formatted_data = split_data(line.strip())
    formatted_string = remove_extra_spaces(formatted_data)
    processed_sequence = remove_00and0x_spaces(formatted_string)
    processed_lines.append(processed_sequence)

# Write processed data to output file
write_multiple_lines_to_out_file(output_filename, processed_lines)

print(f'Output written to {output_filename}')

'''
# Example usage:
input_data = '810a001f01240003016cff0203030f0c0040006519553e44424800003f4901'

# Example usage:
formatted_data = split_data(input_data)
print(f'Formatted data: {formatted_data}')
formatted_string = remove_extra_spaces(formatted_data)
print("formatted_string\n")
print(formatted_string)
processed_sequence = remove_00and0x_spaces(formatted_string)
print("processed_sequence")
print(processed_sequence)

data_lines = [
    processed_sequence
]

filename = 'test.out'
write_multiple_lines_to_out_file(filename, data_lines)
print(f'Output written to {filename}')

'''
