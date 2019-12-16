<template>
<div>
    <a-table
    :columns="columns"
    :dataSource="data"
    bordered
  >
    <template
      slot="code"
      slot-scope="text"
    >
      <span>http://www.rechengparty.com/wx/kanjiahuodong/{{text}}</span>
    </template>
    <template
      slot="title"
      slot-scope="currentPageData"
    >
      热橙派对砍价活动
    </template>
  </a-table>
</div>

</template>
<script>
import axios from "axios";

const columns = [
  {
    title: "邀请链接",
    dataIndex: "code",
    scopedSlots: { customRender: 'code' },
  },
  {
    title: "状态",
    dataIndex: "is_used",
  },
    {
    title: "微信单号",
    dataIndex: "wx_order_id",
  },
];

export default {
  data() {
    return {
      data:[],
      columns,
    };
  },
  mounted() {
    let url = "/wx/unique_pwds";
    axios.get(url).then(res => {
      res = res.data;
      this.data=res.list
      console.log(res.list);
    });
  }
};
</script>
<style>
th.column-money,
td.column-money {
  text-align: right !important;
}
</style>