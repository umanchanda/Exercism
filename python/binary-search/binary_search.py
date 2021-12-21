def find(search_list, value):
    left = 0
    right = len(search_list)-1
    if right >= left:
        mid = left + (right-left) // 2
        if search_list[mid] == value:
            return mid
        elif search_list[mid] > value:
            return find(search_list[left:mid], value)
        else:
            return mid + 1 + find(search_list[mid+1:right+1], value)
    else:
        raise ValueError("not found")
