#Given a string, find the length of the longest substring without repeating characters.
def longestsub(s):
    start = 0
    l = len(s)
    maxi = 1
    count = 1
    visited = []
    for end in range(1,l):
        if s[start] == s[end] or s[end] in visited:
            start = end
            count = 1
            visited = []
        else:
            count +=1
            visited.append(s[end])
            maxi = max(maxi,count)
        # print("char {} and count{} ".format(s[end],count))
    return maxi

s = input("Enter the substring  ")
print("length of the longest substring without repeating characters: ",longestsub(s))