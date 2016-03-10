/*
 * Copyright (C) 2016-2017 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */
'use strict';


var intlData = {
    zh: {
      "title": "串行库",
      "description": "串行库是一个Web服务，加密迹象活泼的Ubuntu模型断言。",
      "home": "家",
      "brand": "牌",
      "brand-description": "设备品牌名称",
      "model": "机型",
      "model-description": "设备型号的名称",
      "revision": "版本",
      "revision-description": "该设备的修订",
      "models": "楷模",
      "models_available": "以下型号可供选择",
      "version": "版本",
      "edit-model": "编辑模型",
      "new-model": "新模式",
      "add-new-model": "添加新模式",
      "save": "保存",
      "cancel": "取消",
      "private-key": "私钥签名",
      "private-key-description": "将用于签署设备标识的签署密钥",
      "public-keys": "公共密钥",

      // Error messages
      "error-nil-data": "未初始化的POST数据",
      "error-sign-empty": "没有签名提供的数据",
      "error-decode-json": "错误解码JSON",
      "error-model-not-found": "无法找到匹配的品牌，型号和修订模型",
      "error-format-assertions": "格式化出错断言",
      "error-read-private-key": "错误读取私钥",
      "error-signing-assertions": "登录错误的断言",
      "error-fetch-models": "错误获取模型",
      "error-invalid-model": "无效的模型ID",
      "error-get-model": "找不到模型",
      "error-model-data": "没有模型数据提供",
      "error-creating-model": "错误创建模型",
      "error-updating-model": "错误更新模型",
      "error-decode-key": "错误解码的base64签名密钥",
      "error-validate-model": "品牌和型号必须提供与修订必须大于零",
      "error-model-exists": "用相同的品牌，型号和版本A设备已经存在",
      "error-invalid-key": "该签名密钥无效",
      "error-created-model": "无法找到创建的模型",
      "error-validate-new-model": "品牌，型号和签名密钥，必须提供与修订必须大于零"
    },

		en: {
      "title": "Serial Vault",
      "description": "The Serial Vault is a web service that cryptographically signs snappy Ubuntu model assertions.",
      "home": "Home",
			"brand": "Brand",
      "brand-description": "The name of the device brand",
      "model": "Model",
      "model-description": "The name of the device model",
      "revision": "Revision",
      "revision-description": "The revision of the device",
      "models": "Models",
      "models_available": "The following models are available",
      "version": "Version",
      "edit-model": "Edit Model",
      "new-model": "New Model",
      "add-new-model": "Add a new model",
      "save": "Save",
      "cancel": "Cancel",
      "private-key": "Private Key for Signing",
      "private-key-description": "The signing-key that will be used to sign the device identity",
      "public-keys": "Public Keys",

      // Error messages
      "error-nil-data": "Uninitialized POST data",
      "error-sign-empty": "No data supplied for signing",
      "error-decode-json": "Error decoding JSON",
      "error-model-not-found": "Cannot find model with the matching brand, model and revision",
      "error-format-assertions": "Error formatting the assertions",
      "error-read-private-key": "Error reading the private key",
      "error-signing-assertions": "Error signing the assertions",
      "error-fetch-models": "Error fetching the models",
      "error-invalid-model": "Invalid model ID",
      "error-get-model": "Cannot find the model",
      "error-model-data": "No model data supplied",
      "error-creating-model": "Error creating the model",
      "error-updating-model": "Error updating the model",
      "error-decode-key": "Error decoding the base64 Signing Key",
      "error-validate-model": "The Brand and Model must be supplied and Revision must be greater than zero",
      "error-model-exists": "A device with the same Brand, Model and Revision already exists",
      "error-invalid-key": "The Signing-key is invalid",
      "error-created-model": "Cannot find the created model",
      "error-validate-new-model": "The Brand, Model and Signing-Key must be supplied and Revision must be greater than zero"
		}
};

module.exports = intlData;