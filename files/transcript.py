transcript1 = open('transcript1.txt', 'r')
transcript2 = open('transcript2.txt', 'r')

# Append the two transcripts into a new file
transcript = open('transcript.txt', 'w')

# Transcript 1 has the first line, transcript 2 has the second line and so on
for line1, line2 in zip(transcript1, transcript2):
    if line1 == "ftqigt ecuvng":
        decoded_line = ""
        for char in line1:
            if char == " ":
                decoded_line += " "
            else:
                decoded_line += chr(ord(char) - 2)

        line1 = decoded_line + "\n"
    transcript.write(line1)
    transcript.write(line2)

transcript1.close()
transcript2.close()
transcript.close()

# Print out the new transcript
transcript = open('transcript.txt', 'r')
for line in transcript:
    print(line, end='')