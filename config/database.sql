create table products(
    kode_produk varchar(10)  not null primary key,
    nama_produk varchar(100) not null,
    stok        smallint,
    harga       integer
);

create table orders(
    id_order serial primary key,
    tanggal_order timestamp,
    payment_method varchar(100),
    total integer
);

create table order_details(
    id_order_detail serial primary key not null,
    id_order integer not null,
    kode_produk varchar(10) not null,
    qty smallint,

    constraint id_order_foreign foreign key (id_order) references orders on update cascade on delete restrict ,
    constraint kode_produk_foreign foreign key (kode_produk) references products on update cascade on delete restrict
)

