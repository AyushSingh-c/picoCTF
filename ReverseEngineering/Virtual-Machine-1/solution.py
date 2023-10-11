import sys
from fractions import Fraction as R
red_axis = int(sys.argv[1])
print(f"{red_axis=}")
diff_1_top = red_axis * R(24, 12) * R(16, 8) * R(16, 8)
diff_1_low = red_axis * R(24, 8) * R(16, 8)
diff_1 = (diff_1_top + diff_1_low) / 2
print(f"{diff_1=}")
diff_2_top = diff_1 * R(24, 12) * R(24, 12) * R(40, 8)
diff_2_low = diff_1 * R(24, 12) * R(24, 8) * R(24, 8)
diff_2 = (diff_2_top + diff_2_low) / 2
print(f"{diff_2=}")
diff_3_top = diff_1 * R(24, 12) * R(16, 8)**5 * R(24, 8)
diff_3_low = diff_2 * R(24, 12) * R(40, 8)
diff_3 = -1 * (diff_3_top + diff_3_low) / 2
print(f"{diff_3=}")
diff_4_top = diff_3 * R(24, 12) * R(16, 8) * R(16, 8)
diff_4_low = diff_3 * R(24, 8) * R(16, 8)
diff_4 = (diff_4_top + diff_4_low) / 2
print(f"{diff_4=}")
blue_axis = float(diff_4)
print(f"{blue_axis=}")