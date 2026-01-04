OUTPUT_DIR=output
INSTALL_DIR=/usr/local/bin

all: clean build

build:
	@echo "正在编译 todo..."
	go build -o $(OUTPUT_DIR)/todo ./cmd/todo
	@echo "正在编译 track..."
	go build -o $(OUTPUT_DIR)/track ./cmd/track
	@echo "编译完成！"

clean:
	@echo "正在清理旧文件..."
	rm -rf $(OUTPUT_DIR)
	mkdir -p $(OUTPUT_DIR)

install: build
	@echo "正在安装到系统 $(INSTALL_DIR)..."
	sudo install -m 755 $(OUTPUT_DIR)/todo $(INSTALL_DIR)/todo
	sudo install -m 755 $(OUTPUT_DIR)/track $(INSTALL_DIR)/track
	@echo "安装成功！"

uninstall:
	@echo "正在卸载..."
	sudo rm -f $(INSTALL_DIR)/todo
	sudo rm -f $(INSTALL_DIR)/track
	@echo "卸载完成！"