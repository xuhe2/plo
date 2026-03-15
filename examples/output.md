
# 🛠 提示词流水线执行技能 (Skill: Plo-Pipeline)

## 第一部分：全局流程拓扑 (Workflow Topology)
> [PROTOCOL]: 此部分仅用于确定逻辑流转，严禁在此执行具体指令。

| 当前节点 ID | 节点摘要 | 跳转条件 (Condition) | 下一个节点 (Target) |
| :--- | :--- | :--- | :--- |
| **Node001** | 检查我的本地文件夹(仓库) |  [IF] "DEFAULT" <br>  |  ➡️ **Node002** <br>  |
| **Node002** | 是否有文件 |  [IF] "是" <br>  [IF] "否" <br>  |  ➡️ **Node003** <br>  ➡️ **Node004** <br>  |
| **Node003** | 检查项目语言 |  [IF] "Unkown" <br>  [IF] "Yes" <br>  |  ➡️ **Node005** <br>  ➡️ **Node006** <br>  |
| **Node004** | 输出: 无项目信息 |  -  |  🏁 END_OF_FLOW  |
| **Node005** | 输出: 无法识别 |  -  |  🏁 END_OF_FLOW  |
| **Node006** | 输出: 项目的编程语言是: XXX |  -  |  🏁 END_OF_FLOW  |

---

## 第二部分：节点指令详情 (Node Specifications)
> [PROTOCOL]: 确定当前节点后，请严格执行以下 [INSTRUCTION] 内容。


### 📍 Node001

**[INSTRUCTION]**
"""
检查我的本地文件夹(仓库)
"""

---

### 📍 Node002

**[INSTRUCTION]**
"""
是否有文件
"""

---

### 📍 Node003

**[INSTRUCTION]**
"""
检查项目语言
"""

---

### 📍 Node004

**[INSTRUCTION]**
"""
输出: 无项目信息
"""

---

### 📍 Node005

**[INSTRUCTION]**
"""
输出: 无法识别
"""

---

### 📍 Node006

**[INSTRUCTION]**
"""
输出: 项目的编程语言是: XXX
"""

---

