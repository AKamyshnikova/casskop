(window.webpackJsonp=window.webpackJsonp||[]).push([[51],{186:function(e,t,a){"use strict";a.r(t),a.d(t,"frontMatter",(function(){return c})),a.d(t,"metadata",(function(){return o})),a.d(t,"rightToc",(function(){return i})),a.d(t,"default",(function(){return l}));var r=a(2),n=a(9),s=(a(0),a(200)),c={id:"3_5_backup_restore",title:"Backup and restore",sidebar_label:"Backup and restore"},o={id:"5_operations/3_5_backup_restore",title:"Backup and restore",description:"In order to provide Backup/Restore abilities we use InstaCluster's cassandra-sidecar project and add it to each Cassandra node to spawn. We want to thant Instaclustr for the modifications they made to make it work with CassKop!",source:"@site/docs/5_operations/3_5_backup_restore.md",permalink:"/casskop/docs/5_operations/3_5_backup_restore",editUrl:"https://github.com/Orange-OpenSource/casskop/edit/master/website/docs/5_operations/3_5_backup_restore.md",sidebar_label:"Backup and restore",sidebar:"docs",previous:{title:"Multi-CassKop",permalink:"/casskop/docs/5_operations/3_multi_casskop"},next:{title:"Upgrade Operator",permalink:"/casskop/docs/5_operations/4_upgrade_operator"}},i=[{value:"Backup",id:"backup",children:[{value:"Supported storage",id:"supported-storage",children:[]},{value:"Life cycle of the CassandraBackup object",id:"life-cycle-of-the-cassandrabackup-object",children:[]}]},{value:"Restore",id:"restore",children:[{value:"Entities",id:"entities",children:[]}]}],p={rightToc:i};function l(e){var t=e.components,a=Object(n.a)(e,["components"]);return Object(s.b)("wrapper",Object(r.a)({},p,a,{components:t,mdxType:"MDXLayout"}),Object(s.b)("p",null,"In order to provide Backup/Restore abilities we use InstaCluster's ",Object(s.b)("a",Object(r.a)({parentName:"p"},{href:"https://github.com/instaclustr/cassandra-sidecar"}),"cassandra-sidecar project")," and add it to each Cassandra node to spawn. We want to thant Instaclustr for the modifications they made to make it work with CassKop!"),Object(s.b)("h2",{id:"backup"},"Backup"),Object(s.b)("p",null,"It is possible to backup keyspaces or tables from cluster managed by Casskop."),Object(s.b)("p",null,"To start or schedule a backup, you create an object of type CassandraBackup:"),Object(s.b)("pre",null,Object(s.b)("code",Object(r.a)({parentName:"pre"},{className:"language-yaml"}),'apiVersion: db.orange.com/v1alpha1\nkind: CassandraBackup\nmetadata:\n  name: nightly-cassandra-backup\n  labels:\n    app: cassandra\nspec:\n  cassandracluster: test-cluster\n  datacenter: dc1\n  storageLocation: s3://cassie\n  snapshotTag: SnapshotTag2\n  secret: cloud-backup-secrets\n  schedule: "@midnight"\n  entities: k1.t1,k2.t3\n')),Object(s.b)("h3",{id:"supported-storage"},"Supported storage"),Object(s.b)("p",null,"The following storage options for storing the backups are:"),Object(s.b)("ul",null,Object(s.b)("li",{parentName:"ul"},"s3 (as in the example above)"),Object(s.b)("li",{parentName:"ul"},"gcp"),Object(s.b)("li",{parentName:"ul"},"azure"),Object(s.b)("li",{parentName:"ul"},"oracle cloud")),Object(s.b)("p",null,"More details can be found on ",Object(s.b)("a",Object(r.a)({parentName:"p"},{href:"https://github.com/instaclustr/cassandra-backup"}),"Instaclustr's Cassandra backup page")),Object(s.b)("h3",{id:"life-cycle-of-the-cassandrabackup-object"},"Life cycle of the CassandraBackup object"),Object(s.b)("p",null,"When this object gets created, CassKop does a few checks to ensure :"),Object(s.b)("ul",null,Object(s.b)("li",{parentName:"ul"},"the specified Cassandra cluster exists"),Object(s.b)("li",{parentName:"ul"},"if there is a secret that it has the expected parameters depending on the chosen backend"),Object(s.b)("li",{parentName:"ul"},"if there is a schedule that its format is correct (crontab like)")),Object(s.b)("p",null,"Then, if all those checks pass, it triggers the backup if there is no schedule, or creates a Cron task with the specified schedule."),Object(s.b)("p",null,"When this object gets deleted, if there is a scheduled task, it is unscheduled."),Object(s.b)("p",null,"When this object gets updated, and the change is located in the spec section, CassKop unschedules the existing task and schedules a new one with the new parameters provided."),Object(s.b)("h2",{id:"restore"},"Restore"),Object(s.b)("p",null,"Following the same logic, a CassandraRestore object must be created to trigger a restore:\nA restore must refer to a Backup object in K8S"),Object(s.b)("pre",null,Object(s.b)("code",Object(r.a)({parentName:"pre"},{className:"language-yaml"}),"apiVersion: db.orange.com/v1alpha1\nkind: CassandraRestore\nmetadata:\n  name: nightly-cassandra-backup\n  labels:\n    app: cassandra\nspec:\n  backup:\n    name: nightly-cassandra-backup\n  cluster: test-cluster\n  entites: k1.t1\n")),Object(s.b)("h3",{id:"entities"},"Entities"),Object(s.b)("p",null,"In the restore phase, you can specify a subset of the entities specified in the backup. You can then backup 2 tables and only restore one."))}l.isMDXComponent=!0},200:function(e,t,a){"use strict";a.d(t,"a",(function(){return u})),a.d(t,"b",(function(){return h}));var r=a(0),n=a.n(r);function s(e,t,a){return t in e?Object.defineProperty(e,t,{value:a,enumerable:!0,configurable:!0,writable:!0}):e[t]=a,e}function c(e,t){var a=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),a.push.apply(a,r)}return a}function o(e){for(var t=1;t<arguments.length;t++){var a=null!=arguments[t]?arguments[t]:{};t%2?c(Object(a),!0).forEach((function(t){s(e,t,a[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(a)):c(Object(a)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(a,t))}))}return e}function i(e,t){if(null==e)return{};var a,r,n=function(e,t){if(null==e)return{};var a,r,n={},s=Object.keys(e);for(r=0;r<s.length;r++)a=s[r],t.indexOf(a)>=0||(n[a]=e[a]);return n}(e,t);if(Object.getOwnPropertySymbols){var s=Object.getOwnPropertySymbols(e);for(r=0;r<s.length;r++)a=s[r],t.indexOf(a)>=0||Object.prototype.propertyIsEnumerable.call(e,a)&&(n[a]=e[a])}return n}var p=n.a.createContext({}),l=function(e){var t=n.a.useContext(p),a=t;return e&&(a="function"==typeof e?e(t):o(o({},t),e)),a},u=function(e){var t=l(e.components);return n.a.createElement(p.Provider,{value:t},e.children)},d={inlineCode:"code",wrapper:function(e){var t=e.children;return n.a.createElement(n.a.Fragment,{},t)}},b=n.a.forwardRef((function(e,t){var a=e.components,r=e.mdxType,s=e.originalType,c=e.parentName,p=i(e,["components","mdxType","originalType","parentName"]),u=l(a),b=r,h=u["".concat(c,".").concat(b)]||u[b]||d[b]||s;return a?n.a.createElement(h,o(o({ref:t},p),{},{components:a})):n.a.createElement(h,o({ref:t},p))}));function h(e,t){var a=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var s=a.length,c=new Array(s);c[0]=b;var o={};for(var i in t)hasOwnProperty.call(t,i)&&(o[i]=t[i]);o.originalType=e,o.mdxType="string"==typeof e?e:r,c[1]=o;for(var p=2;p<s;p++)c[p]=a[p];return n.a.createElement.apply(null,c)}return n.a.createElement.apply(null,a)}b.displayName="MDXCreateElement"}}]);