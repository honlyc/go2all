<template>
    <div class="app-container">
        <div class="filter-container">
            <el-input
                    v-model="listQuery.title"
                    :placeholder="$t('table.title')"
                    style="width: 200px;"
                    class="filter-item"
                    @keyup.enter.native="handleFilter"/>
            <el-input placeholder="请输入内容" v-model="listQuery.type" class="input-with-select">
                <el-select v-model="select" slot="prepend" placeholder="查询类型">
                    {{with .Columns}}
                    {{range $k, $v := .}}
                    <el-option label="{{.Label}}" value="{{$k}}"></el-option>
                    {{end}}
                    {{end}}
                </el-select>
                <el-button slot="append" icon="el-icon-search"></el-button>
            </el-input>
            <el-button class="filter-item" type="primary" icon="el-icon-search" @click="handleFilter">
                搜索
            </el-button>
            <el-button
                    class="filter-item"
                    style="margin-left: 10px;"
                    type="primary"
                    icon="el-icon-edit"
                    @click="handleCreate">
                添加
            </el-button>
        </div>
        <el-table
                :key="tableKey"
                v-loading="listLoading"
                :data="list"
                border
                fit
                highlight-current-row
                style="width: 100%;"
                @sort-change="sortChange"
        >
            {{with .Columns}}
            {{range $k, $v := .}}
            <el-table-column :label="$t('{{.Label}}')" prop="id" align="center">
                <template slot-scope="{row}">
                    <span>{{ToPropName .Key }}</span>
                </template>
            </el-table-column>
            {{end}}
            {{end}}
            <el-table-column :label="$t('table.actions')" align="center" class-name="small-padding fixed-width"
                             width="280">
                <template slot-scope="{row}">
                    <el-button type="" size="mini" @click="handleUpdate(row)">
                        编辑
                    </el-button>
                    <el-button type="danger" size="mini" @click="handleUpdate(row)">
                        删除
                    </el-button>
                </template>
            </el-table-column>
        </el-table>

        <pagination
                v-show="total>0"
                :total="total"
                :page.sync="listQuery.page"
                :limit.sync="listQuery.limit"
                @pagination="getList"/>

        {{with .Columns}}
        {{range $k, $v := .}}
        <el-form-item label-width="150px" label="{{.Label}}">
            <span>{{.Key}}</span>
        </el-form-item>
        {{end}}
        {{end}}
    </div>
</template>
<script>
import {create, fetchList, update} from '@/api/{{.Name}}'
import {parseTime} from '@/utils'
import Pagination from '@/components/Pagination' // secondary package based on el-pagination

export default {
    name: 'ComplexTable',
    components: {Pagination},
    filters: {
        statusFilter(status) {
            const statusMap = {
                published: 'success',
                draft: 'info',
                deleted: 'danger'
            }
            return statusMap[status]
        },
        typeFilter(type) {
            return calendarTypeKeyValue[type]
        }
    },
    data() {
        return {
            tableKey: 0,
            list: null,
            total: 0,
            listLoading: true,
            listQuery: {
                page: 1,
                limit: 20,
                importance: undefined,
                title: undefined,
                type: undefined,
                sort: '+id'
            },
            showReviewer: false,
            temp: {
                id: undefined,
                importance: 1,
                remark: '',
                timestamp: new Date(),
                title: '',
                type: '',
                status: 'published'
            },
            dialogFormVisible: false,
            dialogStatus: '',
            textMap: {
                update: 'Edit',
                create: 'Create'
            },
            dialogPvVisible: false,
            pvData: [],
            rules: {
                type: [{required: true, message: 'type is required', trigger: 'change'}],
                timestamp: [{type: 'date', required: true, message: 'timestamp is required', trigger: 'change'}],
                title: [{required: true, message: 'title is required', trigger: 'blur'}]
            },
            downloadLoading: false
        }
    },
    created() {
        this.getList()
    },
    methods: {
        getList() {
            this.listLoading = true
            fetchList(this.listQuery).then(response => {
                this.list = response.data.items
                this.total = response.data.total
                // Just to simulate the time of the request
                this.listLoading = false
            })
        },
        handleFilter() {
            this.listQuery.page = 1
            this.getList()
        },
        handleModifyStatus(row, status) {
            this.$message({
                message: '操作成功',
                type: 'success'
            })
            row.status = status
        },
        sortChange(data) {
            const {prop, order} = data
            if (prop === 'id') {
                this.sortByID(order)
            }
        },
        sortByID(order) {
            if (order === 'ascending') {
                this.listQuery.sort = '+id'
            } else {
                this.listQuery.sort = '-id'
            }
            this.handleFilter()
        },
        resetTemp() {
            this.temp = {
                id: undefined,
                importance: 1,
                remark: '',
                timestamp: new Date(),
                title: '',
                status: 'published',
                type: ''
            }
        },
        handleCreate() {
            this.resetTemp()
            this.dialogStatus = 'create'
            this.dialogFormVisible = true
            this.$nextTick(() => {
                this.$refs['dataForm'].clearValidate()
            })
        },
        createData() {
            this.$refs['dataForm'].validate((valid) => {
                if (valid) {
                    this.temp.id = parseInt(Math.random() * 100) + 1024 // mock a id
                    this.temp.author = 'vue-element-admin'
                    createCabinet(this.temp).then(() => {
                        this.list.unshift(this.temp)
                        this.dialogFormVisible = false
                        this.$notify({
                            title: '成功',
                            message: '创建成功',
                            type: 'success',
                            duration: 2000
                        })
                    })
                }
            })
        },
        handleUpdate(row) {
            this.temp = Object.assign({}, row) // copy obj
            this.temp.timestamp = new Date(this.temp.timestamp)
            this.dialogStatus = 'update'
            this.dialogFormVisible = true
            this.$nextTick(() => {
                this.$refs['dataForm'].clearValidate()
            })
        },
        updateData() {
            this.$refs['dataForm'].validate((valid) => {
                if (valid) {
                    const tempData = Object.assign({}, this.temp)
                    tempData.timestamp = +new Date(tempData.timestamp) // change Thu Nov 30 2017 16:41:05 GMT+0800 (CST) to 1512031311464
                    updateCabinet(tempData).then(() => {
                        for (const v of this.list) {
                            if (v.id === this.temp.id) {
                                const index = this.list.indexOf(v)
                                this.list.splice(index, 1, this.temp)
                                break
                            }
                        }
                        this.dialogFormVisible = false
                        this.$notify({
                            title: '成功',
                            message: '更新成功',
                            type: 'success',
                            duration: 2000
                        })
                    })
                }
            })
        },
        handleDelete(row) {
            this.$notify({
                title: '成功',
                message: '删除成功',
                type: 'success',
                duration: 2000
            })
            const index = this.list.indexOf(row)
            this.list.splice(index, 1)
        },
        handleDownload() {
            this.downloadLoading = true
        },
        formatJson(filterVal, jsonData) {
            return jsonData.map(v => filterVal.map(j => {
                if (j === 'timestamp') {
                    return parseTime(v[j])
                } else {
                    return v[j]
                }
            }))
        }
    }
}
</script>

