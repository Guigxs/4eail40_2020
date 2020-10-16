class Rectangle:
    def __init__(self, width, height):
        self.width = width
        self.height = height

    def set_width(self, width):
        self.width = width

    def set_height(self, height):
        self.height = height


class Square(Rectangle):
    def __init__(self, size):
        super().__init__(size, size)

    def set_size(self, size):
        self.width = size
        self.height = size

    # What about code calling set_width and set_height ?


if __name__ == "__main__":
    mySquare = Square(10)
    myRectangle = Rectangle(2, 3)
    myRectangles = [mySquare, myRectangle]
    for shape in myRectangles:
        shape.set_height(5)  # wtf!!
