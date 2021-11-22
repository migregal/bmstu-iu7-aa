.PHONY : all run debug runmem gcov clean func check unit

SRC_DIR := ./src
INC_DIR := ./inc
OUT_DIR := ./out

CC := gcc
CFLAGS := -std=gnu17 -I$(INC_DIR) -Wall -Werror -Wfloat-equal -Wfloat-conversion -Wpedantic -pg -D__TICKS_COUNT__
LFLAGS := -lm -lpthread -lc

INCS := $(filter-out $(INC_DIR)/main.h, $(SRCS:$(SRC_DIR)/%.c=$(INC_DIR)/%.h))
SRCS := $(wildcard $(SRC_DIR)/*.c)
OBJS := $(SRCS:$(SRC_DIR)/%.c=$(OUT_DIR)/%.o)

EXEC := app

THREADS=16

all: $(EXEC)

run: $(EXEC)
	./$^

$(EXEC): $(OBJS) $(INCS)
	@$(CC) $(CFLAGS) -o $@ $(OBJS) $(LFLAGS)

$(OUT_DIR):
	@mkdir -p $(OUT_DIR)

${OUT_DIR}/%.o : ${SRC_DIR}/%.c $(OUT_DIR)
	@$(CC) $(CFLAGS) -c $< -o $@

$(OUT_DIR)/check_%.o : $(UTESTS_DIR)/check_%.c $(OUT_DIR)
	@$(CC) $(CFLAGS) -o $@ -c $<

$(INC_DIR)/%.h:

clean:
	@$(RM) -r $(EXEC) $(OUT_DIR)
