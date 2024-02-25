"use client";
import Typography from "@mui/joy/Typography";
import { Card, DonutChart, List, ListItem } from "@tremor/react";

function classNames(...classes: string[]) {
  return classes.filter(Boolean).join(" ");
}

const data = [
  {
    name: "Base Goal",
    amount: 1940,
    share: "32.1%",
    color: "bg-cyan-500",
  },
  {
    name: "Food",
    amount: 900,
    share: "19.6%",
    color: "bg-blue-500",
  },
  {
    name: "Exercise",
    amount: 150,
    share: "18.6%",
    color: "bg-indigo-500",
  },
];

export default function Example() {
  return (
    <>
      <Card className="sm:mx-auto sm:max-w-lg">
        <Typography level="h4" component="h1">
          Calories
        </Typography>
        <Typography level="h5" component="h2">
          Remaining = Goal - Food + Exercise
        </Typography>

        <DonutChart
          className="mt-8"
          data={data}
          category="amount"
          index="name"
          showTooltip={false}
          colors={["cyan", "blue", "indigo", "violet", "fuchsia"]}
          showAnimation={true}
        />
        <p className="mt-8 flex items-center justify-between text-tremor-label text-tremor-content dark:text-dark-tremor-content">
          <span>Category</span>
          <span>Amount / Share</span>
        </p>
        <List className="mt-2">
          {data.map((item) => (
            <ListItem key={item.name} className="space-x-6">
              <div className="flex items-center space-x-2.5 truncate">
                <span
                  className={classNames(
                    item.color,
                    "h-2.5 w-2.5 shrink-0 rounded-sm"
                  )}
                  aria-hidden={true}
                />
                <span className="truncate dark:text-dark-tremor-content-emphasis">
                  {item.name}
                </span>
              </div>
              <div className="flex items-center space-x-2">
                <span className="font-medium tabular-nums text-tremor-content-strong dark:text-dark-tremor-content-strong">
                  {item.amount}
                </span>
                <span className="rounded-tremor-small bg-tremor-background-subtle px-1.5 py-0.5 text-tremor-label font-medium tabular-nums text-tremor-content-emphasis dark:bg-dark-tremor-background-subtle dark:text-dark-tremor-content-emphasis">
                  {item.share}
                </span>
              </div>
            </ListItem>
          ))}
        </List>
      </Card>
    </>
  );
}
