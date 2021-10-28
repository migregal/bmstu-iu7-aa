.PHONY : all run debug runmem gcov clean func check unit

SRC_DIR := ./src
INC_DIR := ./inc
OUT_DIR := ./out
UTESTS_DIR := ./unit_tests
FTESTS_DIR := ./func_tests

CC := gcc
CFLAGS := -std=gnu17 -I$(INC_DIR) -Wall -Werror -Wfloat-equal -Wfloat-conversion -Wpedantic -pg -D__TICKS_COUNT__
LFLAGS := -lm -lpthread -lc

INCS := $(filter-out $(INC_DIR)/main.h, $(SRCS:$(SRC_DIR)/%.c=$(INC_DIR)/%.h))
SRCS := $(wildcard $(SRC_DIR)/*.c)
OBJS := $(SRCS:$(SRC_DIR)/%.c=$(OUT_DIR)/%.o)

EXEC := app.exe

THREADS=16
TYPE=1

all: $(EXEC)

run: $(EXEC)
	./$^ ${THREADS} ${TYPE}

$(EXEC): $(OBJS) $(INCS)
	@$(CC) $(CFLAGS) -o $@ $(OBJS) $(LFLAGS)

debug : CFLAGS += -g -fno-inline -fno-omit-frame-pointer -O0
debug : all

prof : CFLAGS += -O0 -D__FILE_READING__
prof : all

$(OUT_DIR):
	@mkdir -p $(OUT_DIR)

${OUT_DIR}/%.o : ${SRC_DIR}/%.c $(OUT_DIR)
	@$(CC) $(CFLAGS) -c $< -o $@

$(OUT_DIR)/check_%.o : $(UTESTS_DIR)/check_%.c $(OUT_DIR)
	@$(CC) $(CFLAGS) -o $@ -c $<

$(INC_DIR)/%.h:

clean:
	@$(RM) $(EXEC) $(UEXEC) ${OUT_DIR}/*
	@rm -rf $(OUT_DIR)