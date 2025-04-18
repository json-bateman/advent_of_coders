import importlib
import pathlib
import aoc_util

def main():
    times = list()
    for day in range(1, 26):
        try:
            puzzle = importlib.import_module(f"day{day:02}")
        except ModuleNotFoundError:
            continue
        puzzle_input = pathlib.Path(__file__).parent / "input" / f"{day:02}.txt"
        a, b, t = aoc_util.run_solution(puzzle.solve, puzzle_input)
        times.append(t)
        print(f"[Day {day:02}] {t: >8.1f} ms  A:{a} B:{b}")
    print(f"Total time: {sum(times)} ms")
    
if __name__ == "__main__":
    main()
