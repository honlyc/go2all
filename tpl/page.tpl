<template>
  <div class="app-container">
    <div class="filter-container">
      <el-form :inline="true" class="demo-form-inline">
        <el-form-item>
          <el-input
            placeholder="请输入内容"
            v-model="listQuery.value"
            @keyup.enter.native="handleFilter"
            class="input-with-select">
            <el-select v-model="listQuery.select" slot="prepend" placeholder="查询类型" style="width: 120px;">
                {{- range $k, $v := .Columns}}{{- if not .IgFilter}}
                <el-option label="{{.Label}}" value="{{$k}}"></el-option>
                {{- end}}{{end}}
            </el-select>
          </el-input>
        </el-form-item>
        <el-button class="filter-item" type="primary" icon="el-icon-search" @click="handleFilter">
          {{To18IN "table.search"}}
        </el-button>
        <el-button
          class="filter-item"
          style="margin-left: 10px;"
          type="primary"
          icon="el-icon-edit"
          @click="handleCreate">
          {{To18IN "table.add"}}
        </el-button>
      </el-form>
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
        {{- range $k, $v := .Columns}}{{- if not .IgnoreShow}}
        <el-table-column label="{{.Label}}" prop="{{.Key}}" align="center">
            <template slot-scope="{row}">
                <span>{{ToPropName .Key -}}</span>
            </template>
        </el-table-column>
        {{- end}}{{end}}
      <el-table-column :label="$t('table.actions')" align="center" class-name="small-padding fixed-width" width="280">
        <template slot-scope="{row}">
          <el-button type="" size="mini" @click="handleUpdate(row)">
            {{To18IN "table.edit"}}
          </el-button>
          <el-button type="danger" size="mini" @click="handleDelete(row)">
            {{To18IN "table.delete"}}
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

    <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFormVisible">
      <el-form
        ref="dataForm"
        :rules="rules"
        :model="temp"
        label-position="left"
        label-width="100px"
        style="width: 400px; margin-left:50px;">
        {{- range $k, $v := .Columns}}{{- if not .IgnoreEdit}}
        <el-form-item label-width="150px" label="{{.Label}}">
            <el-input v-model="temp.{{.Key}}" {{if not .CanModify}}:disabled="dialogStatus!=='create'"{{end}}/>
        </el-form-item>
        {{- end}}{{end}}
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">
          {{To18IN "table.cancel"}}
        </el-button>
        <el-button type="primary" @click="dialogStatus==='create'?createData():updateData()">
          {{To18IN "table.confirm"}}
        </el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {create, fetchList, remove, update} from '@/api/{{.Name}}'
import Pagination from '@/components/Pagination' // secondary package based on el-pagination

export default {
  name: '{{.MName}}List',
  components: {Pagination},
  filters: {
  },
  data() {
    var self = this
    return {
      tableKey: 0,
      list: null,
      total: 0,
      listLoading: true,
      listQuery: {
        page_num: 1,
        page_size: 10,
        {{- range $k, $v := .Columns}}
        {{.Key}}: undefined,
        {{- end}}
        sort: '+id'
      },
      temp: {
        {{- range $k, $v := .Columns}}{{- if not .IgnoreEdit}}
        {{.Key}}: '',
        {{- end}}{{end}}
      },
      dialogFormVisible: false,
      dialogStatus: '',
      textMap: {
        update: 'Edit',
        create: 'Create'
      },
      rules: {
        type: [{required: true, message: 'type is required', trigger: 'change'}],
        timestamp: [{type: 'date', required: true, message: 'timestamp is required', trigger: 'change'}],
        username: [{required: true, message: 'title is required', trigger: 'blur'}],
      }
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
        this.total = response.data.total_size
        this.listLoading = false
      }).catch(err => {
        this.listLoading = false
      })
    },
    handleFilter() {
      this.listQuery.page = 1
      this.getList()
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
        {{- range $k, $v := .Columns}}{{- if not .IgnoreEdit}}
        {{.Key}}: '',
        {{- end}}{{end}}
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
          create(this.temp).then(() => {
            this.list.unshift(this.temp)
            this.dialogFormVisible = false
            this.$notify({
              title: '成功',
              message: '创建成功',
              type: 'success',
              duration: 2000
            })
            this.getList();
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
          update(tempData).then(() => {
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
      var that = this
      this.$confirm('此操作将删除 ' + row.username, '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        callback: function(a, b) {
          if (a == 'confirm') { // 确认后再执行
            remove(row.id).then(res => {
              console.log(res)
              that.$message({
                message: res.data.msg,
                type: 'success'
              })
              that.$notify({
                title: '成功',
                message: '删除成功',
                type: 'success',
                duration: 2000
              })
              const index = that.list.indexOf(row)
              that.list.splice(index, 1)
            }).catch(err => {
              console.error(err)
              that.$message({
                message: err,
                type: 'error'
              })
            })
          }
        }
      })
    }
  }
}
</script>
