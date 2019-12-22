<template>
  <div>
    <a-table :columns="columns" :dataSource="data" bordered>
      <template slot="code" slot-scope="text">
        <span>http://www.rechengparty.com/wx/kanjiahuodong/{{ text }}</span>
      </template>
      <template slot="wx_order_id" slot-scope="text">
        <span
          style="cursor:pointer;display:block;width:100%;height:100%;"
          @click="showModal(text)"
          >{{ text }}</span
        >
      </template>
      <template slot="title" slot-scope="currentPageData">
        热橙派对砍价活动
      </template>
    </a-table>

    <a-modal title="Basic Modal" v-model="visible" width="750px" @ok="handleOk">
      <div v-if="order.length">
        店名：{{ order[0].store_name }}
        价格：{{ order[0].old_price }}
      </div>
      <div>
        <a-input placeholder="店名" allowClear v-model="storeName" />
        <div style="margin-bottom:8px;"></div>
        <a-input placeholder="价格" allowClear v-model="price" />
        <div style="margin-bottom:8px;"></div>
        <a-input placeholder="钥匙" allowClear v-model="key" />

        <div style="margin-top:8px;margin-bottom:16px;"></div>
        <a-button type="primary" shape="round" block @click="submit"
          >提交</a-button
        >
      </div>
      <h3>砍价参与者(微信账号信息)</h3>
      <a-list itemLayout="horizontal" :dataSource="joiner">
        <a-list-item slot="renderItem" slot-scope="item, index">
          <a-list-item-meta>
            <a-avatar slot="avatar" :src="item.wx_user_image" />
          </a-list-item-meta>
          <a-list-item-meta>
            <a slot="title">{{ item.wx_user_name }}</a>
          </a-list-item-meta>
          <a-list-item-meta>
            <a slot="title">{{ item.wx_addr }}</a>
          </a-list-item-meta>
          <a-list-item-meta>
            <a slot="title">{{ (item.wx_sex = 1 ? "男" : "女") }}</a>
          </a-list-item-meta>
          <a-list-item-meta>
            <a slot="title">砍价{{ item.help_amount }}元</a>
          </a-list-item-meta>
          <a-list-item-meta>
            <a slot="title">{{ format(item.create_at * 1000) }}</a>
          </a-list-item-meta>
        </a-list-item>
      </a-list>
    </a-modal>
  </div>
</template>
<script>
import axios from "axios";
import moment from "moment";

const columns = [
  {
    title: "邀请链接",
    dataIndex: "code",
    scopedSlots: { customRender: "code" }
  },
  {
    title: "状态",
    dataIndex: "is_used"
  },
  {
    title: "微信单号",
    dataIndex: "wx_order_id",
    scopedSlots: { customRender: "wx_order_id" }
  }
];

export default {
  data() {
    return {
      visible: false,
      data: [],
      columns,
      joiner: [],
      order: [],
      orderId: "",
      storeName: "",
      price: "",
      key: ""
    };
  },
  mounted() {
    console.log(moment);
    let url = "/wx/unique_pwds";
    axios.get(url).then(res => {
      res = res.data;
      this.data = res.list;
      console.log(res.list);
    });
  },
  methods: {
    showModal(text) {
      console.log(text);
      this.orderId = text;
      this.visible = true;
      let url = "/wx/admin_info?id=" + text;
      axios.get(url).then(res => {
        this.joiner = res.data.list2;
        this.order = res.data.list1;
        console.log(res);
      });
    },
    handleOk(e) {
      console.log(e);
      this.visible = false;
    },
    format(time) {
      return moment(time).format("YYYY年MM月DD日 hh时mm分ss秒");
    },
    submit() {
      console.log(this.orderId);
      console.log(this.value);
      if (this.key == "" || this.price == "" || this.storeName == "") {
        this.$message.success("你的输入有问题，请检查", 5);
        return;
      }
      let value = {
        key: this.key,
        price: this.price,
        storeName: this.storeName,
        orderId:this.orderId,
      };
      let url = "/wx/postadmin";
      axios.post(url, value).then(res => {
        console.log(res.data);
      });
      this.handleOk()
    }
  }
};
</script>
<style>
th.column-money,
td.column-money {
  text-align: right !important;
}
</style>
