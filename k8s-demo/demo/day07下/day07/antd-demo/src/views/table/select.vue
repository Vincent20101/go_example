<template>
    <div style="margin:20px;">
        <a-button type="primary">批量删除</a-button>
        <!-- columns定义表格的列数据 -->
        <!-- dataSource定义表格的源数据列表 -->
        <!-- loading一般在发起list请求时置为true，请求结束后置为false -->
        <a-table
            :dataSource="dataSource"
            :columns="columns"
            :loading="false"
            size="small"
            row-key="name"
            :rowSelection="{selectedRowKeys: selectedRowKeys, onChange: onSelectChange}">
            <!-- row-key用于自定义selectedRowKeys数组中的元素 -->
            <!-- selectedRowKeys决定当前选择了哪些行，默认的元素是 dataSource 中定义的 key -->
            <!-- onChange 是个回调函数，用于表格选择框发生变化时执行的函数 -->

            <!-- 自定义每一列的数据 -->
            <!-- column表示columns列数据 -->
            <!-- record表示dataSource的数据 -->
            <template #bodyCell="{column, record}">
                <template v-if="column.dataIndex == 'name'">
                    <span style="font-weight:bold;color:blue;">{{ record.name }} - {{ record.age }}</span>
                </template>
            </template>
            <!-- 扩展,展开某一行数据 -->
            <template #expandedRowRender="{ record }">
                <p style="margin:0px;">
                    {{ record.address }}123
                </p>
            </template>
        </a-table>
    </div>
</template>

<script>
import {ref} from 'vue'
export default({
    setup() {
        //list数据
        const dataSource = ref([
            {
                key: '1',
                name: '胡彦斌',
                age: 32,
                address: '西湖区湖底公园1号'
            },
            {
                key: '2',
                name: '胡彦祖',
                age: 42,
                address: '西湖区湖底公园1号'
            }
        ])
        const columns = ref([
            // title 列名字
            // dataIndex 列数据在数据项中对应的路径，支持通过数组查询嵌套路径
            // key Vue需要的key，如果已经设置了唯一的dataIndex，可以忽略这个属性
            {
                title: '姓名',
                dataIndex: 'name',
                key: 'name'
            },
            {
                title: '年龄',
                dataIndex: 'age',
                key: 'age'
            },
            {
                title: '住址',
                dataIndex: 'address',
                key: 'address'
            }
        ])
        const selectedRowKeys = ref([])
        const onSelectChange = key => {
            console.log(key)
            selectedRowKeys.value = key
        }
        return {
            dataSource,
            columns,
            selectedRowKeys,
            onSelectChange
        }
    },
})
</script>
