
# 🛠 提示词流水线执行技能 (Skill: Plo-Pipeline)

## 第一部分：全局流程拓扑 (Workflow Topology)
> [PROTOCOL]: 此部分仅用于确定逻辑流转，严禁在此执行具体指令。

| 当前节点 ID | 节点名称 | 跳转条件 (Condition) | 下一个节点 (Target) |
| :--- | :--- | :--- | :--- |
| **Node001** | 检查项目语言 |  [IF] "Yes" <br>  [IF] "Unkown" <br>  |  ➡️ **Node003** (输出: 项目的编程语言是: XXX) <br>  ➡️ **Node002** (输出: 无法识别) <br>  |
| **Node002** | 输出: 无法识别 |  -  |  🏁 END_OF_FLOW  |
| **Node003** | 输出: 项目的编程语言是: XXX |  -  |  🏁 END_OF_FLOW  |
| **Node004** | 检查我的本地文件夹(仓库) |  [IF] "DEFAULT" <br>  |  ➡️ **Node005** (是否有文件) <br>  |
| **Node005** | 是否有文件 |  [IF] "否" <br>  [IF] "是" <br>  |  ➡️ **Node006** (输出: 无项目信息) <br>  ➡️ **Node001** (检查项目语言) <br>  |
| **Node006** | 输出: 无项目信息 |  -  |  🏁 END_OF_FLOW  |

---

## 第二部分：节点指令详情 (Node Specifications)
> [PROTOCOL]: 确定当前节点后，请严格执行以下 [INSTRUCTION] 内容。


### 📍 Node001 | 检查项目语言

**[INSTRUCTION]**
"""
检查项目语言
"""

**[CONTEXT/DATA]**
> 检查项目语言

---

### 📍 Node002 | 输出: 无法识别

**[INSTRUCTION]**
"""
输出: 无法识别
"""

**[CONTEXT/DATA]**
> 输出: 无法识别

---

### 📍 Node003 | 输出: 项目的编程语言是: XXX

**[INSTRUCTION]**
"""
输出: 项目的编程语言是: XXX
"""

**[CONTEXT/DATA]**
> 输出: 项目的编程语言是: XXX

---

### 📍 Node004 | 检查我的本地文件夹(仓库)

**[INSTRUCTION]**
"""
检查我的本地文件夹(仓库)
"""

**[CONTEXT/DATA]**
> 检查我的本地文件夹(仓库)

---

### 📍 Node005 | 是否有文件

**[INSTRUCTION]**
"""
是否有文件
"""

**[CONTEXT/DATA]**
> 是否有文件

---

### 📍 Node006 | 输出: 无项目信息

**[INSTRUCTION]**
"""
输出: 无项目信息
"""

**[CONTEXT/DATA]**
> 输出: 无项目信息

---

