### Jump

You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position. Return true if you can reach the last index, or false otherwise.

### Two city scheduling

A recruiter plans to hire n people and conducts their interviews at two different locations of the company. He evaluates the cost of inviting candidates to both these locations. The plan is to invite 50% at one location, and the rest at the other location, keeping costs to a minimum.

### Assign cookies

You are given two integer arrays:

- **greedFactors**: Represents the minimum cookie size required for each child to be content. Specifically, `greedFactors[i]` denotes the minimum cookie size that child `i` will accept to be satisfied.
- **cookieSizes**: Represents the sizes of the available cookies. For example, `cookieSizes[j]` denotes the size of cookie `j`.

Your objective is to distribute the cookies in a way that maximizes the number of content children. A child `i` is content if they receive a cookie that is at least as large as their greed factor, i.e., `cookieSizes[j] >= greedFactors[i]`. Each child can be assigned at most one cookie, and each cookie can only be given to one child.