#!/usr/bin/env python3
import os
import pathlib

import humanize
import matplotlib
import pandas
import pandas as pd


@matplotlib.ticker.FuncFormatter
def bytes_major_formatter(x, pos):
    return humanize.naturalsize(int(x), binary=True).replace(".0", "")


def render(df: pandas.DataFrame, out: os.PathLike, max_power_of_two: int, max_step: int):
    plot = df.plot(x="Step", y=["VmRSS", "VmStk"])

    plot.set_ylabel('Memory')
    plot.set_yscale("log")
    powers_of_two = [2**j for j in range(1, max_power_of_two+1)]
    yticks = powers_of_two[::2]
    plot.set_yticks(yticks)
    plot.yaxis.set_major_formatter(bytes_major_formatter)
    plot.set_ylim(top=powers_of_two[-1])

    plot.set_xlabel('1 step == 10 microseconds')
    plot.set_xlim(left=0, right=max_step)
    plot.legend(loc=4)

    fig = plot.get_figure()
    fig.tight_layout()
    fig.savefig(out)


def main():
    df_go = pd.read_csv('/tmp/stackoverflow-go/go.csv')
    df_cgo = pd.read_csv('/tmp/stackoverflow-go/cgo.csv')
    max_power_of_two = 29
    max_step = max(df_go["Step"].max(), df_cgo["Step"].max())

    render(df_go, pathlib.Path('/tmp/stackoverflow-go/go.png'), max_power_of_two, max_step)
    render(df_cgo, pathlib.Path('/tmp/stackoverflow-go/cgo.png'), max_power_of_two, max_step)


if __name__ == "__main__":
    main()
