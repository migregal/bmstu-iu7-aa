import random

def gen_matr_to_file(fin, n, m):
    fin.write(str(n) + " " + str(m) + "\n")
    for i in range(n):
        fin.write(" ".join([str(random.randint(-10_000, 10_000)) for k in range(m)]) + "\n")

def main():
    for i in [2**5 * 2**t for t in range(8)]:
        with open("./data/" + str(i) + "x" + str(i) + ".txt", "w") as fin:
            print(str(i) + "x" + str(i))

            gen_matr_to_file(fin, i, i)
            gen_matr_to_file(fin, i, i)


if __name__ == "__main__":
    main()
