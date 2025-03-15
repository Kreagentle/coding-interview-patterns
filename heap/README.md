### Largest Number After Digit Swaps by Parity

You are given a positive integer num. You can swap any two digits of num as long as they share the same parity (both are odd or both are even).
Your task is to return the largest possible value of num after performing any number of such swaps.

### Find right interval

You are given an array of intervals, where intervals[i] = [starti, endi] and each starti is unique.
The right interval for an interval i is an interval j such that startj >= endi and startj is minimized. Note that i may equal j.
Return an array of right interval indices for each interval i. If no right interval exists for interval i, then put -1 at index i.
